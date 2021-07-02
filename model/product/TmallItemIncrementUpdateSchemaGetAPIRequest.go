package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemIncrementUpdateSchemaGetAPIRequest 天猫增量更新商品规则获取 API请求
// tmall.item.increment.update.schema.get
//
// 增量方式修改天猫商品的规则获取的API。<br/>1.接口返回支持增量修改的字段以及相应字段的规则。<br/>2.如果入参xml_data指定了更新的字段，则只返回指定字段的规则（ISV如果功能性很强，如明确更新Title，请拼装好次字段以提升API整体性能）；<br/>3.ISV初次接入，开发阶段，此字段不填可以看到所有支持增量的字段；但是如果上线功能明确，请尽量遵守第2条<br/>4.如果ISV对字段规则非常清晰，可以直接组装入参数据提交到tmall.item.schema.increment.update进行数据更新。但是最好不要写死，比如每天还是有对此接口功能的一次比对。<br/>---（感谢爱慕旗舰店提供API命名）
type TmallItemIncrementUpdateSchemaGetAPIRequest struct {
	model.Params
	// 需要编辑的商品ID
	_itemId int64
	// 如果入参xml_data指定了更新的字段，则只返回指定字段的规则（ISV如果功能性很强，如明确更新Title，请拼装好此字段以提升API整体性能）
	_xmlData string
}

// NewTmallItemIncrementUpdateSchemaGetRequest 初始化TmallItemIncrementUpdateSchemaGetAPIRequest对象
func NewTmallItemIncrementUpdateSchemaGetRequest() *TmallItemIncrementUpdateSchemaGetAPIRequest {
	return &TmallItemIncrementUpdateSchemaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemIncrementUpdateSchemaGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.increment.update.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemIncrementUpdateSchemaGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 需要编辑的商品ID
func (r *TmallItemIncrementUpdateSchemaGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TmallItemIncrementUpdateSchemaGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is XmlData Setter
// 如果入参xml_data指定了更新的字段，则只返回指定字段的规则（ISV如果功能性很强，如明确更新Title，请拼装好此字段以提升API整体性能）
func (r *TmallItemIncrementUpdateSchemaGetAPIRequest) SetXmlData(_xmlData string) error {
	r._xmlData = _xmlData
	r.Set("xml_data", _xmlData)
	return nil
}

// Get XmlData Getter
func (r TmallItemIncrementUpdateSchemaGetAPIRequest) GetXmlData() string {
	return r._xmlData
}
