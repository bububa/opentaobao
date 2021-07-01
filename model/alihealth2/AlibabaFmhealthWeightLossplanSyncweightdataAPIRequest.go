package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest
薄荷同步三方记录 API请求
alibaba.fmhealth.weight.lossplan.syncweightdata

用于三方薄荷同步数据到健康会员 */
type AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest struct {
	model.Params
	// 阿里健康id
	_tpUserId int64
	// 记录体重
	_weight string
	// 记录日期
	_recordDate string
}

// New
