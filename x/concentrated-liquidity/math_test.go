package concentrated_liquidity_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	cl "github.com/osmosis-labs/osmosis/v12/x/concentrated-liquidity"
)

func (suite *KeeperTestSuite) TestGetLiquidityFromAmounts() {
	testCases := map[string]struct {
		currentSqrtP      sdk.Dec
		sqrtPHigh         sdk.Dec
		sqrtPLow          sdk.Dec
		amount0Desired    sdk.Int
		amount1Desired    sdk.Int
		expectedLiquidity string
	}{
		"happy path": {
			currentSqrtP:      sdk.MustNewDecFromStr("70.710678118654752440"), // 5000
			sqrtPHigh:         sdk.MustNewDecFromStr("74.161984870956629487"), // 5500
			sqrtPLow:          sdk.MustNewDecFromStr("67.416615162732695594"), // 4545
			amount0Desired:    sdk.NewInt(1000000),
			amount1Desired:    sdk.NewInt(5000000000),
			expectedLiquidity: "1517882343.751510418088349649",
		},
	}

	for name, tc := range testCases {
		tc := tc

		suite.Run(name, func() {
			liquidity := cl.GetLiquidityFromAmounts(tc.currentSqrtP, tc.sqrtPLow, tc.sqrtPHigh, tc.amount0Desired, tc.amount1Desired)
			suite.Require().Equal(tc.expectedLiquidity, liquidity.String())
			liq0 := cl.Liquidity0(tc.amount0Desired, tc.currentSqrtP, tc.sqrtPHigh)
			liq1 := cl.Liquidity1(tc.amount1Desired, tc.currentSqrtP, tc.sqrtPLow)
			liq := sdk.MinDec(liq0, liq1)
			suite.Require().Equal(liq.String(), liquidity.String())

		})
	}
}

// liquidity1 takes an amount of asset1 in the pool as well as the sqrtpCur and the nextPrice
// sqrtPriceA is the smaller of sqrtpCur and the nextPrice
// sqrtPriceB is the larger of sqrtpCur and the nextPrice
// liquidity1 = amount1 / (sqrtPriceB - sqrtPriceA)
func (suite *KeeperTestSuite) TestLiquidity1() {
	testCases := map[string]struct {
		currentSqrtP      sdk.Dec
		sqrtPLow          sdk.Dec
		amount1Desired    sdk.Int
		expectedLiquidity string
	}{
		"happy path": {
			currentSqrtP:      sdk.MustNewDecFromStr("70.710678118654752440"), // 5000
			sqrtPLow:          sdk.MustNewDecFromStr("67.416615162732695594"), // 4545
			amount1Desired:    sdk.NewInt(5000000000),
			expectedLiquidity: "1517882343.751510418088349649",
			// https://www.wolframalpha.com/input?i=5000000000+%2F+%2870.710678118654752440+-+67.416615162732695594%29
		},
	}

	for name, tc := range testCases {
		tc := tc

		suite.Run(name, func() {
			liquidity := cl.Liquidity1(tc.amount1Desired, tc.currentSqrtP, tc.sqrtPLow)
			suite.Require().Equal(tc.expectedLiquidity, liquidity.String())
		})
	}
}

// TestLiquidity0 tests that liquidity0 takes an amount of asset0 in the pool as well as the sqrtpCur and the nextPrice
// sqrtPriceA is the smaller of sqrtpCur and the nextPrice
// sqrtPriceB is the larger of sqrtpCur and the nextPrice
// liquidity0 = amount0 * (sqrtPriceA * sqrtPriceB) / (sqrtPriceB - sqrtPriceA)
func (suite *KeeperTestSuite) TestLiquidity0() {
	testCases := map[string]struct {
		currentSqrtP      sdk.Dec
		sqrtPHigh         sdk.Dec
		amount0Desired    sdk.Int
		expectedLiquidity string
	}{
		"happy path": {
			currentSqrtP:      sdk.MustNewDecFromStr("70.710678118654752440"), // 5000
			sqrtPHigh:         sdk.MustNewDecFromStr("74.161984870956629487"), // 5500
			amount0Desired:    sdk.NewInt(1000000),
			expectedLiquidity: "1519437308.014768571721000000", // TODO: should be 1519437308.014768571720923239
			// https://www.wolframalpha.com/input?i=1000000+*+%2870.710678118654752440*+74.161984870956629487%29+%2F+%2874.161984870956629487+-+70.710678118654752440%29
		},
	}

	for name, tc := range testCases {
		tc := tc

		suite.Run(name, func() {
			liquidity := cl.Liquidity0(tc.amount0Desired, tc.currentSqrtP, tc.sqrtPHigh)
			suite.Require().Equal(tc.expectedLiquidity, liquidity.String())
		})
	}
}

// TestGetNextSqrtPriceFromAmount0RoundingUp tests that getNextSqrtPriceFromAmount0RoundingUp utilizes
// the current squareRootPrice, liquidity of denom0, and amount of denom0 that still needs
// to be swapped in order to determine the next squareRootPrice
// PATH 1
// if (amountRemaining * sqrtPriceCurrent) / amountRemaining  == sqrtPriceCurrent AND (liquidity * 2) + (amountRemaining * sqrtPriceCurrent) >= (liquidity * 2)
// sqrtPriceNext = (liquidity * 2 * sqrtPriceCurrent) / ((liquidity * 2) + (amountRemaining * sqrtPriceCurrent))
// PATH 2
// else
// sqrtPriceNext = ((liquidity * 2)) / (((liquidity * 2) / (sqrtPriceCurrent)) + (amountRemaining))
func (suite *KeeperTestSuite) TestGetNextSqrtPriceFromAmount0RoundingUp() {
	testCases := map[string]struct {
		liquidity             sdk.Dec
		sqrtPCurrent          sdk.Dec
		amount0Remaining      sdk.Dec
		sqrtPriceNextExpected string
	}{
		"happy path": {
			liquidity:             sdk.MustNewDecFromStr("1517882343.751510418088349649"), // liquidity0 calculated above
			sqrtPCurrent:          sdk.MustNewDecFromStr("70.710678118654752440"),
			amount0Remaining:      sdk.NewDec(133700),
			sqrtPriceNextExpected: "70.491153655973103127",
			// https://www.wolframalpha.com/input?i=%281517882343.751510418088349649+*+2+*+70.710678118654752440%29+%2F+%28%281517882343.751510418088349649+*+2%29+%2B+%28133700+*+70.710678118654752440%29%29
		},
	}

	for name, tc := range testCases {
		tc := tc

		suite.Run(name, func() {
			sqrtPriceNext := cl.GetNextSqrtPriceFromAmount0RoundingUp(tc.sqrtPCurrent, tc.liquidity, tc.amount0Remaining)
			suite.Require().Equal(tc.sqrtPriceNextExpected, sqrtPriceNext.String())
		})
	}
}

// TestGetNextSqrtPriceFromAmount1RoundingDown tests that getNextSqrtPriceFromAmount1RoundingDown
// utilizes the current squareRootPrice, liquidity of denom1, and amount of denom1 that still needs
// to be swapped in order to determine the next squareRootPrice
// sqrtPriceNext = sqrtPriceCurrent + (amount1Remaining / liquidity1)
func (suite *KeeperTestSuite) TestGetNextSqrtPriceFromAmount1RoundingDown() {
	testCases := map[string]struct {
		liquidity             sdk.Dec
		sqrtPCurrent          sdk.Dec
		amount1Remaining      sdk.Dec
		sqrtPriceNextExpected string
	}{
		"happy path": {
			liquidity:             sdk.MustNewDecFromStr("1519437308.014768571721000000"), // liquidity1 calculated above
			sqrtPCurrent:          sdk.MustNewDecFromStr("70.710678118654752440"),         // 5000000000
			amount1Remaining:      sdk.NewDec(42000000),
			sqrtPriceNextExpected: "70.738319930382329008",
			// https://www.wolframalpha.com/input?i=70.710678118654752440+%2B++++%2842000000+%2F+1519437308.014768571721000000%29
		},
	}

	for name, tc := range testCases {
		tc := tc

		suite.Run(name, func() {
			sqrtPriceNext := cl.GetNextSqrtPriceFromAmount1RoundingDown(tc.sqrtPCurrent, tc.liquidity, tc.amount1Remaining)
			suite.Require().Equal(tc.sqrtPriceNextExpected, sqrtPriceNext.String())
		})
	}
}

// TestCalcAmount0Delta tests that calcAmount0 takes the asset with the smaller liquidity in the pool as well as the sqrtpCur and the nextPrice and calculates the amount of asset 0
// sqrtPriceA is the smaller of sqrtpCur and the nextPrice
// sqrtPriceB is the larger of sqrtpCur and the nextPrice
// calcAmount0Delta = (liquidity * (sqrtPriceB - sqrtPriceA)) / (sqrtPriceB * sqrtPriceA)
func (suite *KeeperTestSuite) TestCalcAmount0Delta() {
	testCases := map[string]struct {
		liquidity       sdk.Dec
		sqrtPCurrent    sdk.Dec
		sqrtPUpper      sdk.Dec
		amount0Expected string
	}{
		"happy path": {
			liquidity:       sdk.MustNewDecFromStr("1517882343.751510418088349649"), // we use the smaller liquidity between liq0 and liq1
			sqrtPCurrent:    sdk.MustNewDecFromStr("70.710678118654752440"),         // 5000
			sqrtPUpper:      sdk.MustNewDecFromStr("74.161984870956629487"),         // 5500
			amount0Expected: "998976.618347426747968399",                            // TODO: should be 998976.618347426388356630
			// https://www.wolframalpha.com/input?i=%281517882343.751510418088349649+*+%2874.161984870956629487+-+70.710678118654752440+%29%29+%2F+%2870.710678118654752440+*+74.161984870956629487%29
		},
	}

	for name, tc := range testCases {
		tc := tc

		suite.Run(name, func() {
			amount0 := cl.CalcAmount0Delta(tc.liquidity, tc.sqrtPCurrent, tc.sqrtPUpper, false)
			suite.Require().Equal(tc.amount0Expected, amount0.String())
		})
	}
}

// TestCalcAmount1Delta tests that calcAmount1 takes the asset with the smaller liquidity in the pool as well as the sqrtpCur and the nextPrice and calculates the amount of asset 1
// sqrtPriceA is the smaller of sqrtpCur and the nextPrice
// sqrtPriceB is the larger of sqrtpCur and the nextPrice
// calcAmount1Delta = liq * (sqrtPriceB - sqrtPriceA)
func (suite *KeeperTestSuite) TestCalcAmount1Delta() {
	testCases := map[string]struct {
		liquidity       sdk.Dec
		sqrtPCurrent    sdk.Dec
		sqrtPLower      sdk.Dec
		amount1Expected string
	}{
		"happy path": {
			liquidity:       sdk.MustNewDecFromStr("1517882343.751510418088349649"), // we use the smaller liquidity between liq0 and liq1
			sqrtPCurrent:    sdk.MustNewDecFromStr("70.710678118654752440"),         // 5000
			sqrtPLower:      sdk.MustNewDecFromStr("67.416615162732695594"),         // 4545
			amount1Expected: "5000000000.000000000000000000",
			// https://www.wolframalpha.com/input?i=1517882343.751510418088349649+*+%2870.710678118654752440+-+67.416615162732695594%29
		},
	}

	for name, tc := range testCases {
		tc := tc

		suite.Run(name, func() {
			amount1 := cl.CalcAmount1Delta(tc.liquidity, tc.sqrtPCurrent, tc.sqrtPLower, false)
			suite.Require().Equal(tc.amount1Expected, amount1.String())
		})
	}
}

func (suite *KeeperTestSuite) TestComputeSwapState() {
	testCases := map[string]struct {
		sqrtPCurrent          sdk.Dec
		sqrtPTarget           sdk.Dec
		liquidity             sdk.Dec
		amountRemaining       sdk.Dec
		zeroForOne            bool
		expectedSqrtPriceNext string
		expectedAmountIn      string
		expectedAmountOut     string
	}{
		"happy path: trade asset0 for asset1": {
			sqrtPCurrent:          sdk.MustNewDecFromStr("70.710678118654752440"), // 5000
			sqrtPTarget:           sdk.MustNewDecFromStr("70.666662070529219856"), // 4993.777128190373086350
			liquidity:             sdk.MustNewDecFromStr("1517818840.967515822610790519"),
			amountRemaining:       sdk.NewDec(13370),
			zeroForOne:            true,
			expectedSqrtPriceNext: "70.666662070529219856",
			expectedAmountIn:      "13369.999999903622360944",
			expectedAmountOut:     "66808387.149866264039333362",
		},
		"happy path: trade asset1 for asset0": {
			sqrtPCurrent:          sdk.MustNewDecFromStr("70.710678118654752440"), // 5000
			sqrtPTarget:           sdk.MustNewDecFromStr("70.738349405152439867"), // 5003.91407656543054317
			liquidity:             sdk.MustNewDecFromStr("1517818840.967515822610790519"),
			amountRemaining:       sdk.NewDec(42000000),
			zeroForOne:            false,
			expectedSqrtPriceNext: "70.738349405152439867",
			expectedAmountIn:      "42000000.000000000650233591",
			expectedAmountOut:     "8396.714104746015980302",
		},
	}

	for name, tc := range testCases {
		tc := tc

		suite.Run(name, func() {
			sqrtPriceNext, amountIn, amountOut := cl.ComputeSwapStep(tc.sqrtPCurrent, tc.sqrtPTarget, tc.liquidity, tc.amountRemaining, tc.zeroForOne)
			suite.Require().Equal(tc.expectedSqrtPriceNext, sqrtPriceNext.String())
			suite.Require().Equal(tc.expectedAmountIn, amountIn.String())
			suite.Require().Equal(tc.expectedAmountOut, amountOut.String())
		})
	}
}
