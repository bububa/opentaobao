package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpRangeAddAPIResponse 添加活动范围 API返回值
// taobao.ump.range.add
//
// 指定某项活动中，某个商家的某些类型物品（指定商品或者别的）参加或者不参加活动。当活动详情的参与类型为部分或者部分不参加的时候，需要指定具体哪部分参加或者不参加，使用本接口完成操作。比如部分商品满就送，这里的range用来指定具体哪些商品参加满就送活动。
type TaobaoUmpRangeAddAPIResponse struct {
	model.CommonResponse
	TaobaoUmpRangeAddAPIResponseModel
}

// TaobaoUmpRangeAddAPIResponseModel is 添加活动范围 成功返回结果
type TaobaoUmpRangeAddAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_range_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
