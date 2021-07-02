package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayWordpackageUpdateAPIRequest 批量更新词包 API请求
// taobao.subway.wordpackage.update
//
// 批量更新词包
type TaobaoSubwayWordpackageUpdateAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
	// 词包列表
	_wordPackageDTOS []ItemWordPackageDto
}

// NewTaobaoSubwayWordpackageUpdateRequest 初始化TaobaoSubwayWordpackageUpdateAPIRequest对象
func NewTaobaoSubwayWordpackageUpdateRequest() *TaobaoSubwayWordpackageUpdateAPIRequest {
	return &TaobaoSubwayWordpackageUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayWordpackageUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.subway.wordpackage.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayWordpackageUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSubwayWordpackageUpdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSubwayWordpackageUpdateAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaoSubwayWordpackageUpdateAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSubwayWordpackageUpdateAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetWordPackageDTOS is WordPackageDTOS Setter
// 词包列表
func (r *TaobaoSubwayWordpackageUpdateAPIRequest) SetWordPackageDTOS(_wordPackageDTOS []ItemWordPackageDto) error {
	r._wordPackageDTOS = _wordPackageDTOS
	r.Set("word_package_d_t_o_s", _wordPackageDTOS)
	return nil
}

// GetWordPackageDTOS WordPackageDTOS Getter
func (r TaobaoSubwayWordpackageUpdateAPIRequest) GetWordPackageDTOS() []ItemWordPackageDto {
	return r._wordPackageDTOS
}
