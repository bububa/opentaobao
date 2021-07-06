package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItFareBatchaddAPIRequest 【国际机票自有政策】批量添加 API请求
// taobao.alitrip.it.fare.batchadd
//
// 支持自有政策和销售规则批量添加，支持携程的数据格式。淘宝格式为list [object] to json string，object的属性和单条接口一致。每个接入方最多同时只能有1个处理中的导入任务，超过后直接返回失败。文件一定要zip压缩，压缩后大小不超过5M，编码格式utf-8
type TaobaoAlitripItFareBatchaddAPIRequest struct {
	model.Params
	// json格式的字符串，扩展属性，预留
	_extendAttributes string
	// 新增类型，1 自有政策单程 2 自有政策往返 3 销售规则
	_addType int64
	// 文本zip压缩后的数据字节流
	_bytes *model.File
	// 数据格式类型，1 淘宝 2 携程
	_dataType int64
}

// NewTaobaoAlitripItFareBatchaddRequest 初始化TaobaoAlitripItFareBatchaddAPIRequest对象
func NewTaobaoAlitripItFareBatchaddRequest() *TaobaoAlitripItFareBatchaddAPIRequest {
	return &TaobaoAlitripItFareBatchaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItFareBatchaddAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.fare.batchadd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItFareBatchaddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetExtendAttributes is ExtendAttributes Setter
// json格式的字符串，扩展属性，预留
func (r *TaobaoAlitripItFareBatchaddAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extendAttributes", _extendAttributes)
	return nil
}

// GetExtendAttributes ExtendAttributes Getter
func (r TaobaoAlitripItFareBatchaddAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// SetAddType is AddType Setter
// 新增类型，1 自有政策单程 2 自有政策往返 3 销售规则
func (r *TaobaoAlitripItFareBatchaddAPIRequest) SetAddType(_addType int64) error {
	r._addType = _addType
	r.Set("addType", _addType)
	return nil
}

// GetAddType AddType Getter
func (r TaobaoAlitripItFareBatchaddAPIRequest) GetAddType() int64 {
	return r._addType
}

// SetBytes is Bytes Setter
// 文本zip压缩后的数据字节流
func (r *TaobaoAlitripItFareBatchaddAPIRequest) SetBytes(_bytes *model.File) error {
	r._bytes = _bytes
	r.Set("bytes", _bytes)
	return nil
}

// GetBytes Bytes Getter
func (r TaobaoAlitripItFareBatchaddAPIRequest) GetBytes() *model.File {
	return r._bytes
}

// SetDataType is DataType Setter
// 数据格式类型，1 淘宝 2 携程
func (r *TaobaoAlitripItFareBatchaddAPIRequest) SetDataType(_dataType int64) error {
	r._dataType = _dataType
	r.Set("dataType", _dataType)
	return nil
}

// GetDataType DataType Getter
func (r TaobaoAlitripItFareBatchaddAPIRequest) GetDataType() int64 {
	return r._dataType
}
