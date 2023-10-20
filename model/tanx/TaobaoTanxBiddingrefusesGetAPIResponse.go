package tanx

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotanxbiddingrefusesgetAPIResponse tanx竞价失败反馈api API返回值
// taobao.tanx.biddingrefuses.get
//
// 竞价失败反馈根据创意id查询API提供
type TaobaotanxbiddingrefusesgetAPIResponse struct {
	model.CommonResponse
	TaobaotanxbiddingrefusesgetAPIResponseModel
}

// TaobaotanxbiddingrefusesgetAPIResponseModel is tanx竞价失败反馈api 成功返回结果
type TaobaotanxbiddingrefusesgetAPIResponseModel struct {
	XMLName xml.Name `xml:"tanx_biddingrefuses_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回竞价失败对象信息
	BiddingRefuseList []BiddingRefuseDto `json:"bidding_refuse_list,omitempty" xml:"bidding_refuse_list>bidding_refuse_dto,omitempty"`
	// 返回是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
