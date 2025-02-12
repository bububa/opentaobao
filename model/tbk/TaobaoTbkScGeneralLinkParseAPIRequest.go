package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScGeneralLinkParseAPIRequest 淘宝客-服务商-万能解析 API请求
// taobao.tbk.sc.general.link.parse
//
// 淘宝客-推广者-万能转链
type TaobaoTbkScGeneralLinkParseAPIRequest struct {
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

// NewTaobaoTbkScGeneralLinkParseRequest 初始化TaobaoTbkScGeneralLinkParseAPIRequest对象
func NewTaobaoTbkScGeneralLinkParseRequest() *TaobaoTbkScGeneralLinkParseAPIRequest {
	return &TaobaoTbkScGeneralLinkParseAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkScGeneralLinkParseAPIRequest) Reset() {
	r._materialDto = r._materialDto[:0]
	r._relationId = ""
	r._fields = ""
	r._adzoneId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScGeneralLinkParseAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.general.link.parse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScGeneralLinkParseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScGeneralLinkParseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMaterialDto is MaterialDto Setter
// 链接/口令转链
func (r *TaobaoTbkScGeneralLinkParseAPIRequest) SetMaterialDto(_materialDto []LkMaterialDto) error {
	r._materialDto = _materialDto
	r.Set("material_dto", _materialDto)
	return nil
}

// GetMaterialDto MaterialDto Getter
func (r TaobaoTbkScGeneralLinkParseAPIRequest) GetMaterialDto() []LkMaterialDto {
	return r._materialDto
}

// SetRelationId is RelationId Setter
// 渠道管理ID（如是主站选品推广场景，必须入参该字段，且bizSceneId字段需入参2-消费者比价场景，否则二次转链失败）
func (r *TaobaoTbkScGeneralLinkParseAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoTbkScGeneralLinkParseAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetFields is Fields Setter
// 有origin_pid使用场景需入参
func (r *TaobaoTbkScGeneralLinkParseAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTbkScGeneralLinkParseAPIRequest) GetFields() string {
	return r._fields
}

// SetAdzoneId is AdzoneId Setter
// 推广位id，mm_xx_xx_xx pid三段式中的第三段
func (r *TaobaoTbkScGeneralLinkParseAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkScGeneralLinkParseAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

var poolTaobaoTbkScGeneralLinkParseAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkScGeneralLinkParseRequest()
	},
}

// GetTaobaoTbkScGeneralLinkParseRequest 从 sync.Pool 获取 TaobaoTbkScGeneralLinkParseAPIRequest
func GetTaobaoTbkScGeneralLinkParseAPIRequest() *TaobaoTbkScGeneralLinkParseAPIRequest {
	return poolTaobaoTbkScGeneralLinkParseAPIRequest.Get().(*TaobaoTbkScGeneralLinkParseAPIRequest)
}

// ReleaseTaobaoTbkScGeneralLinkParseAPIRequest 将 TaobaoTbkScGeneralLinkParseAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkScGeneralLinkParseAPIRequest(v *TaobaoTbkScGeneralLinkParseAPIRequest) {
	v.Reset()
	poolTaobaoTbkScGeneralLinkParseAPIRequest.Put(v)
}
