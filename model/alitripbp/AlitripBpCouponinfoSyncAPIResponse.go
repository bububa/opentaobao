package alitripbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBpCouponinfoSyncAPIResponse 飞猪广告券信息同步接口 API返回值
// alitrip.bp.couponinfo.sync
//
// 飞猪商业化券信息同步
type AlitripBpCouponinfoSyncAPIResponse struct {
	model.CommonResponse
	AlitripBpCouponinfoSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBpCouponinfoSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBpCouponinfoSyncAPIResponseModel).Reset()
}

// AlitripBpCouponinfoSyncAPIResponseModel is 飞猪广告券信息同步接口 成功返回结果
type AlitripBpCouponinfoSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_bp_couponinfo_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AdResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBpCouponinfoSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBpCouponinfoSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBpCouponinfoSyncAPIResponse)
	},
}

// GetAlitripBpCouponinfoSyncAPIResponse 从 sync.Pool 获取 AlitripBpCouponinfoSyncAPIResponse
func GetAlitripBpCouponinfoSyncAPIResponse() *AlitripBpCouponinfoSyncAPIResponse {
	return poolAlitripBpCouponinfoSyncAPIResponse.Get().(*AlitripBpCouponinfoSyncAPIResponse)
}

// ReleaseAlitripBpCouponinfoSyncAPIResponse 将 AlitripBpCouponinfoSyncAPIResponse 保存到 sync.Pool
func ReleaseAlitripBpCouponinfoSyncAPIResponse(v *AlitripBpCouponinfoSyncAPIResponse) {
	v.Reset()
	poolAlitripBpCouponinfoSyncAPIResponse.Put(v)
}
