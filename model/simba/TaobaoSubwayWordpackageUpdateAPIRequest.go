package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwaywordpackageupdateAPIRequest 批量更新词包 API请求
// taobao.subway.wordpackage.update
//
// 批量更新词包
type TaobaosubwaywordpackageupdateAPIRequest struct {
	model.Params
	// 词包列表
	_wordPackageDTOS []ItemWordPackageDto
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaosubwaywordpackageupdateRequest 初始化TaobaosubwaywordpackageupdateAPIRequest对象
func NewTaobaosubwaywordpackageupdateRequest() *TaobaosubwaywordpackageupdateAPIRequest {
	return &TaobaosubwaywordpackageupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubwaywordpackageupdateAPIRequest) GetApiMethodName() string {
	return "taobao.subway.wordpackage.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubwaywordpackageupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubwaywordpackageupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWordPackageDTOS is WordPackageDTOS Setter
// 词包列表
func (r *TaobaosubwaywordpackageupdateAPIRequest) SetWordPackageDTOS(_wordPackageDTOS []ItemWordPackageDto) error {
	r._wordPackageDTOS = _wordPackageDTOS
	r.Set("word_package_d_t_o_s", _wordPackageDTOS)
	return nil
}

// GetWordPackageDTOS WordPackageDTOS Getter
func (r TaobaosubwaywordpackageupdateAPIRequest) GetWordPackageDTOS() []ItemWordPackageDto {
	return r._wordPackageDTOS
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosubwaywordpackageupdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosubwaywordpackageupdateAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaosubwaywordpackageupdateAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosubwaywordpackageupdateAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
