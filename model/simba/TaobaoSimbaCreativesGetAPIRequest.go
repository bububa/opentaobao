package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacreativesgetAPIRequest 批量获得创意 API请求
// taobao.simba.creatives.get
//
// 取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；&lt;br/&gt;如果同时提供了推广组Id和创意id列表，则优先使用推广组Id；
type TaobaosimbacreativesgetAPIRequest struct {
	model.Params
	// 创意Id数组，最多200个
	_creativeIds []string
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaosimbacreativesgetRequest 初始化TaobaosimbacreativesgetAPIRequest对象
func NewTaobaosimbacreativesgetRequest() *TaobaosimbacreativesgetAPIRequest {
	return &TaobaosimbacreativesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacreativesgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.creatives.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacreativesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacreativesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreativeIds is CreativeIds Setter
// 创意Id数组，最多200个
func (r *TaobaosimbacreativesgetAPIRequest) SetCreativeIds(_creativeIds []string) error {
	r._creativeIds = _creativeIds
	r.Set("creative_ids", _creativeIds)
	return nil
}

// GetCreativeIds CreativeIds Getter
func (r TaobaosimbacreativesgetAPIRequest) GetCreativeIds() []string {
	return r._creativeIds
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbacreativesgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbacreativesgetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaosimbacreativesgetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbacreativesgetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
