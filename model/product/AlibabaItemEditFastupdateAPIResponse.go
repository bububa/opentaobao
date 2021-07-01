package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItemEditFastupdateAPIResponse
商品编辑增量更新 API返回值
alibaba.item.edit.fastupdate

商品编辑增量更新;
<br/>该接口编辑sku，只能更新价格、库存等信息，不能新增sku;
<br/>新增sku用全量接口alibaba.item.edit.submit，先设置销售属性; */
type AlibabaItemEditFastupdateAPIResponse struct {
	model.CommonResponse
	AlibabaItemEditFastupdateAPIResponseModel
}

// AlibabaItemEditFastupdateAPIResponseModel is 商品编辑增量更新 成功返回结果
type AlibabaItemEditFastupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_edit_fastupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品更新时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品所属市场
	Market string `json:"market,omitempty" xml:"market,omitempty"`
}
