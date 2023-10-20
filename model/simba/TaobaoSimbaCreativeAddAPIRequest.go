package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCreativeAddAPIRequest 增加创意 API请求
// taobao.simba.creative.add
//
// 创建一个创意
type TaobaoSimbaCreativeAddAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 创意标题，最多30个汉字
	_title string
	// 创意图片地址，必须是推广组对应商品的图片之一
	_imgUrl string
	// 广审批准文号
	_adExaminationCode string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaoSimbaCreativeAddRequest 初始化TaobaoSimbaCreativeAddAPIRequest对象
func NewTaobaoSimbaCreativeAddRequest() *TaobaoSimbaCreativeAddAPIRequest {
	return &TaobaoSimbaCreativeAddAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCreativeAddAPIRequest) Reset() {
	r._nick = ""
	r._title = ""
	r._imgUrl = ""
	r._adExaminationCode = ""
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCreativeAddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.creative.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCreativeAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCreativeAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCreativeAddAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCreativeAddAPIRequest) GetNick() string {
	return r._nick
}

// SetTitle is Title Setter
// 创意标题，最多30个汉字
func (r *TaobaoSimbaCreativeAddAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoSimbaCreativeAddAPIRequest) GetTitle() string {
	return r._title
}

// SetImgUrl is ImgUrl Setter
// 创意图片地址，必须是推广组对应商品的图片之一
func (r *TaobaoSimbaCreativeAddAPIRequest) SetImgUrl(_imgUrl string) error {
	r._imgUrl = _imgUrl
	r.Set("img_url", _imgUrl)
	return nil
}

// GetImgUrl ImgUrl Getter
func (r TaobaoSimbaCreativeAddAPIRequest) GetImgUrl() string {
	return r._imgUrl
}

// SetAdExaminationCode is AdExaminationCode Setter
// 广审批准文号
func (r *TaobaoSimbaCreativeAddAPIRequest) SetAdExaminationCode(_adExaminationCode string) error {
	r._adExaminationCode = _adExaminationCode
	r.Set("ad_examination_code", _adExaminationCode)
	return nil
}

// GetAdExaminationCode AdExaminationCode Getter
func (r TaobaoSimbaCreativeAddAPIRequest) GetAdExaminationCode() string {
	return r._adExaminationCode
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaCreativeAddAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaCreativeAddAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaCreativeAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCreativeAddRequest()
	},
}

// GetTaobaoSimbaCreativeAddRequest 从 sync.Pool 获取 TaobaoSimbaCreativeAddAPIRequest
func GetTaobaoSimbaCreativeAddAPIRequest() *TaobaoSimbaCreativeAddAPIRequest {
	return poolTaobaoSimbaCreativeAddAPIRequest.Get().(*TaobaoSimbaCreativeAddAPIRequest)
}

// ReleaseTaobaoSimbaCreativeAddAPIRequest 将 TaobaoSimbaCreativeAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCreativeAddAPIRequest(v *TaobaoSimbaCreativeAddAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCreativeAddAPIRequest.Put(v)
}
