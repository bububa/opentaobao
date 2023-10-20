package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCreativesGetAPIRequest 批量获得创意 API请求
// taobao.simba.creatives.get
//
// 取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；&lt;br/&gt;如果同时提供了推广组Id和创意id列表，则优先使用推广组Id；
type TaobaoSimbaCreativesGetAPIRequest struct {
	model.Params
	// 创意Id数组，最多200个
	_creativeIds []string
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaoSimbaCreativesGetRequest 初始化TaobaoSimbaCreativesGetAPIRequest对象
func NewTaobaoSimbaCreativesGetRequest() *TaobaoSimbaCreativesGetAPIRequest {
	return &TaobaoSimbaCreativesGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCreativesGetAPIRequest) Reset() {
	r._creativeIds = r._creativeIds[:0]
	r._nick = ""
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCreativesGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.creatives.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCreativesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCreativesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreativeIds is CreativeIds Setter
// 创意Id数组，最多200个
func (r *TaobaoSimbaCreativesGetAPIRequest) SetCreativeIds(_creativeIds []string) error {
	r._creativeIds = _creativeIds
	r.Set("creative_ids", _creativeIds)
	return nil
}

// GetCreativeIds CreativeIds Getter
func (r TaobaoSimbaCreativesGetAPIRequest) GetCreativeIds() []string {
	return r._creativeIds
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCreativesGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCreativesGetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaCreativesGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaCreativesGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaCreativesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCreativesGetRequest()
	},
}

// GetTaobaoSimbaCreativesGetRequest 从 sync.Pool 获取 TaobaoSimbaCreativesGetAPIRequest
func GetTaobaoSimbaCreativesGetAPIRequest() *TaobaoSimbaCreativesGetAPIRequest {
	return poolTaobaoSimbaCreativesGetAPIRequest.Get().(*TaobaoSimbaCreativesGetAPIRequest)
}

// ReleaseTaobaoSimbaCreativesGetAPIRequest 将 TaobaoSimbaCreativesGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCreativesGetAPIRequest(v *TaobaoSimbaCreativesGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCreativesGetAPIRequest.Put(v)
}
