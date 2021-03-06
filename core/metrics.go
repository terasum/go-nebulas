// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package core

import (
	metrics "github.com/rcrowley/go-metrics"
)

// Metrics for core
var (
	// block metrics
	metricsBlockHeightGauge      = metrics.GetOrRegisterGauge("neb.block.height", nil)
	metricsBlocktailHashGauge    = metrics.GetOrRegisterGauge("neb.block.tailhash", nil)
	metricsBlockRevertTimesGauge = metrics.GetOrRegisterGauge("neb.block.revertcount", nil)
	metricsBlockRevertMeter      = metrics.GetOrRegisterMeter("neb.block.revert", nil)
	metricsBlockOnchainTimer     = metrics.GetOrRegisterTimer("neb.block.onchain", nil)
	metricsTxOnchainTimer        = metrics.GetOrRegisterTimer("neb.tx.onchain", nil)

	// block_pool metrics
	metricsDuplicatedBlock    = metrics.GetOrRegisterCounter("neb.block.duplicated", nil)
	metricsInvalidBlock       = metrics.GetOrRegisterCounter("neb.block.invalid", nil)
	metricsBlockExecutedTimer = metrics.GetOrRegisterTimer("neb.block.executed", nil)
	metricsTxExecutedTimer    = metrics.GetOrRegisterTimer("neb.tx.executed", nil)

	// txpool metrics
	metricsInvalidTx           = metrics.GetOrRegisterCounter("txpool.invalid", nil)
	metricsDuplicateTx         = metrics.GetOrRegisterCounter("txpool.duplicate", nil)
	metricsTxPoolBelowGasPrice = metrics.GetOrRegisterCounter("txpool.below_gas_price", nil)
	metricsTxPoolOutOfGasLimit = metrics.GetOrRegisterCounter("txpool.out_of_gas_limit", nil)

	// transaction metrics
	metricsTxSubmit     = metrics.GetOrRegisterMeter("tx.submit", nil)
	metricsTxExecute    = metrics.GetOrRegisterMeter("tx.execute", nil)
	metricsTxExeSuccess = metrics.GetOrRegisterMeter("tx.execute.success", nil)
	metricsTxExeFailed  = metrics.GetOrRegisterMeter("tx.execute.failed", nil)
)
