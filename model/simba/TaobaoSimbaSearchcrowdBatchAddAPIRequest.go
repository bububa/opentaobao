package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasearchcrowdbatchaddAPIRequest 推广单元增加搜索人群 API请求
// taobao.simba.searchcrowd.batch.add
//
// 推广单元新增搜索人群
type TaobaosimbasearchcrowdbatchaddAPIRequest struct {
	model.Params
	// 被操作者的淘宝昵称
	_nick string
	// 新增人群信息,批量接口,入参为list,溢价(discount)范围为[105,400]
	_adgroupTargetingTags string
	// 推广单元id
	_adgroupId int64
}

// NewTaobaosimbasearchcrowdbatchaddRequest 初始化TaobaosimbasearchcrowdbatchaddAPIRequest对象
func NewTaobaosimbasearchcrowdbatchaddRequest() *TaobaosimbasearchcrowdbatchaddAPIRequest {
	return &TaobaosimbasearchcrowdbatchaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbasearchcrowdbatchaddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.searchcrowd.batch.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbasearchcrowdbatchaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbasearchcrowdbatchaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaosimbasearchcrowdbatchaddAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbasearchcrowdbatchaddAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupTargetingTags is AdgroupTargetingTags Setter
// 新增人群信息,批量接口,入参为list,溢价(discount)范围为[105,400]
func (r *TaobaosimbasearchcrowdbatchaddAPIRequest) SetAdgroupTargetingTags(_adgroupTargetingTags string) error {
	r._adgroupTargetingTags = _adgroupTargetingTags
	r.Set("adgroup_targeting_tags", _adgroupTargetingTags)
	return nil
}

// GetAdgroupTargetingTags AdgroupTargetingTags Getter
func (r TaobaosimbasearchcrowdbatchaddAPIRequest) GetAdgroupTargetingTags() string {
	return r._adgroupTargetingTags
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaosimbasearchcrowdbatchaddAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbasearchcrowdbatchaddAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
