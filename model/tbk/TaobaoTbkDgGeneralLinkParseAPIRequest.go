package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgGeneralLinkParseAPIRequest 淘宝客-推广者-万能解析 API请求
// taobao.tbk.dg.general.link.parse
//
// 淘宝客-推广者-万能转链
type TaobaoTbkDgGeneralLinkParseAPIRequest struct {
	model.Params
	// 链接/口令转链
	_materialDto []LkMaterialDto
	// 渠道管理ID（如是主站选品推广场景，必须入参该字段，且bizSceneId字段需入参2-消费者比价场景，否则二次转链失败）
	_relationId string
	// 有origin_pid使用场景需入参
	_fields string
	// 推广位id，mm_xx_xx_xx pid三段式中的第三段
	_adzoneId int64
}

// NewTaobaoTbkDgGeneralLinkParseRequest 初始化TaobaoTbkDgGeneralLinkParseAPIRequest对象
func NewTaobaoTbkDgGeneralLinkParseRequest() *TaobaoTbkDgGeneralLinkParseAPIRequest {
	return &TaobaoTbkDgGeneralLinkParseAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkDgGeneralLinkParseAPIRequest) Reset() {
	r._materialDto = r._materialDto[:0]
	r._relationId = ""
	r._fields = ""
	r._adzoneId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgGeneralLinkParseAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.general.link.parse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgGeneralLinkParseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkDgGeneralLinkParseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMaterialDto is MaterialDto Setter
// 链接/口令转链
func (r *TaobaoTbkDgGeneralLinkParseAPIRequest) SetMaterialDto(_materialDto []LkMaterialDto) error {
	r._materialDto = _materialDto
	r.Set("material_dto", _materialDto)
	return nil
}

// GetMaterialDto MaterialDto Getter
func (r TaobaoTbkDgGeneralLinkParseAPIRequest) GetMaterialDto() []LkMaterialDto {
	return r._materialDto
}

// SetRelationId is RelationId Setter
// 渠道管理ID（如是主站选品推广场景，必须入参该字段，且bizSceneId字段需入参2-消费者比价场景，否则二次转链失败）
func (r *TaobaoTbkDgGeneralLinkParseAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoTbkDgGeneralLinkParseAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetFields is Fields Setter
// 有origin_pid使用场景需入参
func (r *TaobaoTbkDgGeneralLinkParseAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTbkDgGeneralLinkParseAPIRequest) GetFields() string {
	return r._fields
}

// SetAdzoneId is AdzoneId Setter
// 推广位id，mm_xx_xx_xx pid三段式中的第三段
func (r *TaobaoTbkDgGeneralLinkParseAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkDgGeneralLinkParseAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

var poolTaobaoTbkDgGeneralLinkParseAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkDgGeneralLinkParseRequest()
	},
}

// GetTaobaoTbkDgGeneralLinkParseRequest 从 sync.Pool 获取 TaobaoTbkDgGeneralLinkParseAPIRequest
func GetTaobaoTbkDgGeneralLinkParseAPIRequest() *TaobaoTbkDgGeneralLinkParseAPIRequest {
	return poolTaobaoTbkDgGeneralLinkParseAPIRequest.Get().(*TaobaoTbkDgGeneralLinkParseAPIRequest)
}

// ReleaseTaobaoTbkDgGeneralLinkParseAPIRequest 将 TaobaoTbkDgGeneralLinkParseAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkDgGeneralLinkParseAPIRequest(v *TaobaoTbkDgGeneralLinkParseAPIRequest) {
	v.Reset()
	poolTaobaoTbkDgGeneralLinkParseAPIRequest.Put(v)
}
