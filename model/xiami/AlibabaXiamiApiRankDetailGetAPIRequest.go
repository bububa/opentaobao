package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiRankDetailGetAPIRequest
排行榜详情 API请求
alibaba.xiami.api.rank.detail.get

虾米排行榜详情数据 */
type AlibabaXiamiApiRankDetailGetAPIRequest struct {
	model.Params
	// 榜单ID
	_billboardId int64
	// 调用来源
	_bizCode string
}

// NewAlibabaXiamiApiRankDetailGetRequest 初始化AlibabaXiamiApiRankDetailGetAPIRequest对象
func NewAlibabaXiamiApiRankDetailGetRequest() *AlibabaXiamiApiRankDetailGetAPIRequest {
	return &AlibabaXiamiApiRankDetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiRankDetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.xiami.api.rank.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiRankDetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BillboardId Setter
// 榜单ID
func (r *AlibabaXiamiApiRankDetailGetAPIRequest) SetBillboardId(_billboardId int64) error {
	r._billboardId = _billboardId
	r.Set("billboard_id", _billboardId)
	return nil
}

// Get BillboardId Getter
func (r AlibabaXiamiApiRankDetailGetAPIRequest) GetBillboardId() int64 {
	return r._billboardId
}

// Set is BizCode Setter
// 调用来源
func (r *AlibabaXiamiApiRankDetailGetAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// Get BizCode Getter
func (r AlibabaXiamiApiRankDetailGetAPIRequest) GetBizCode() string {
	return r._bizCode
}
