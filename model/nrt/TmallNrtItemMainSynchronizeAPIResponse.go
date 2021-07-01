package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtItemMainSynchronizeAPIResponse
家装新零售主商品同步至阿里 API返回值
tmall.nrt.item.main.synchronize

同步红星美凯龙存量商品到阿里 */
type TmallNrtItemMainSynchronizeAPIResponse struct {
	model.CommonResponse
	TmallNrtItemMainSynchronizeAPIResponseModel
}

// TmallNrtItemMainSynchronizeAPIResponseModel is 家装新零售主商品同步至阿里 成功返回结果
type TmallNrtItemMainSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_item_main_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	TmallNrtItemMainSynchronize *TmallNrtItemMainSynchronizeResultDo `json:"tmall_nrt_item_main_synchronize,omitempty" xml:"tmall_nrt_item_main_synchronize,omitempty"`
}
