package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayWordpackageUpdateAPIRequest 批量更新词包 API请求
// taobao.subway.wordpackage.update
//
// 批量更新词包
type TaobaoSubwayWordpackageUpdateAPIRequest struct {
	model.Params
	// 词包列表
	_wordPackageDTOS []ItemWordPackageDto
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaoSubwayWordpackageUpdateRequest 初始化TaobaoSubwayWordpackageUpdateAPIRequest对象
func NewTaobaoSubwayWordpackageUpdateRequest() *TaobaoSubwayWordpackageUpdateAPIRequest {
	return &TaobaoSubwayWordpackageUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubwayWordpackageUpdateAPIRequest) Reset() {
	r._wordPackageDTOS = r._wordPackageDTOS[:0]
	r._nick = ""
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayWordpackageUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.subway.wordpackage.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayWordpackageUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayWordpackageUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoSubwayWordpackageUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubwayWordpackageUpdateRequest()
	},
}

// GetTaobaoSubwayWordpackageUpdateRequest 从 sync.Pool 获取 TaobaoSubwayWordpackageUpdateAPIRequest
func GetTaobaoSubwayWordpackageUpdateAPIRequest() *TaobaoSubwayWordpackageUpdateAPIRequest {
	return poolTaobaoSubwayWordpackageUpdateAPIRequest.Get().(*TaobaoSubwayWordpackageUpdateAPIRequest)
}

// ReleaseTaobaoSubwayWordpackageUpdateAPIRequest 将 TaobaoSubwayWordpackageUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubwayWordpackageUpdateAPIRequest(v *TaobaoSubwayWordpackageUpdateAPIRequest) {
	v.Reset()
	poolTaobaoSubwayWordpackageUpdateAPIRequest.Put(v)
}
