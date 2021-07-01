package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractLotteryactivityRegisterAPIRequest
回传抽奖相关参数 API请求
alibaba.interact.lotteryactivity.register

提供接口供三方应用将数据回传到平台 */
type AlibabaInteractLotteryactivityRegisterAPIRequest struct {
	model.Params
	// 入参
	_paramTopUpdateActivityLotteryInfoParam *TopUpdateActivityLotteryInfoParam
}

// New
