package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkItemInfoTemporaryGetAPIRequest 淘宝客商品详情查询（简版）（临时接口） API请求
// taobao.tbk.item.info.temporary.get
//
// 淘宝客商品详情查询（简版）（临时接口）
type TaobaoTbkItemInfoTemporaryGetAPIRequest struct {
	model.Params
	// 商品ID串，用,分割，最大40个
	_numIids string
	// ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
	_ip string
	// 链接形式：1：PC，2：无线，默认：１
	_platform int64
}

// NewTaobaoTbkItemInfoTemporaryGetRequest 初始化TaobaoTbkItemInfoTemporaryGetAPIRequest对象
func NewTaobaoTbkItemInfoTemporaryGetRequest() *TaobaoTbkItemInfoTemporaryGetAPIRequest {
	return &TaobaoTbkItemInfoTemporaryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkItemInfoTemporaryGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.item.info.temporary.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkItemInfoTemporaryGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNumIids is NumIids Setter
// 商品ID串，用,分割，最大40个
func (r *TaobaoTbkItemInfoTemporaryGetAPIRequest) SetNumIids(_numIids string) error {
	r._numIids = _numIids
	r.Set("num_iids", _numIids)
	return nil
}

// GetNumIids NumIids Getter
func (r TaobaoTbkItemInfoTemporaryGetAPIRequest) GetNumIids() string {
	return r._numIids
}

// SetIp is Ip Setter
// ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
func (r *TaobaoTbkItemInfoTemporaryGetAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r TaobaoTbkItemInfoTemporaryGetAPIRequest) GetIp() string {
	return r._ip
}

// SetPlatform is Platform Setter
// 链接形式：1：PC，2：无线，默认：１
func (r *TaobaoTbkItemInfoTemporaryGetAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaoTbkItemInfoTemporaryGetAPIRequest) GetPlatform() int64 {
	return r._platform
}
