package hotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelInfoListGetForHelloAPIResponse 哈罗获取酒店详情 API返回值
// taobao.xhotel.info.list.get.for.hello
//
// 哈罗合作项目，供哈罗合作方批量和增量两种场景下查询已开通城市下的标准酒店及房型信息，不涉及用户登录信息
type TaobaoXhotelInfoListGetForHelloAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelInfoListGetForHelloAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelInfoListGetForHelloAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelInfoListGetForHelloAPIResponseModel).Reset()
}

// TaobaoXhotelInfoListGetForHelloAPIResponseModel is 哈罗获取酒店详情 成功返回结果
type TaobaoXhotelInfoListGetForHelloAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_info_list_get_for_hello_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoXhotelInfoListGetForHelloResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelInfoListGetForHelloAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelInfoListGetForHelloAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelInfoListGetForHelloAPIResponse)
	},
}

// GetTaobaoXhotelInfoListGetForHelloAPIResponse 从 sync.Pool 获取 TaobaoXhotelInfoListGetForHelloAPIResponse
func GetTaobaoXhotelInfoListGetForHelloAPIResponse() *TaobaoXhotelInfoListGetForHelloAPIResponse {
	return poolTaobaoXhotelInfoListGetForHelloAPIResponse.Get().(*TaobaoXhotelInfoListGetForHelloAPIResponse)
}

// ReleaseTaobaoXhotelInfoListGetForHelloAPIResponse 将 TaobaoXhotelInfoListGetForHelloAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelInfoListGetForHelloAPIResponse(v *TaobaoXhotelInfoListGetForHelloAPIResponse) {
	v.Reset()
	poolTaobaoXhotelInfoListGetForHelloAPIResponse.Put(v)
}
