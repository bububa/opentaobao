package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopicturegetAPIRequest 获取图片信息 API请求
// taobao.picture.get
//
// 获取图片信息
type TaobaopicturegetAPIRequest struct {
	model.Params
	// 是否删除,unfroze代表没有删除
	_deleted string
	// 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
	_modifiedTime string
	// 图片标题,最大长度50字符,中英文都算一字符
	_title string
	// 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
	_startDate string
	// 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
	_endDate string
	// 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
	_clientType string
	// 图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc
	_orderBy string
	// 图片url查询接口
	_urls string
	// 图片ID
	_pictureId int64
	// 图片分类ID
	_pictureCategoryId int64
	// 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1
	_pageNo int64
	// 每页条数.每页返回最多返回100条,默认值40
	_pageSize int64
	// 是否获取https的链接
	_isHttps bool
}

// NewTaobaopicturegetRequest 初始化TaobaopicturegetAPIRequest对象
func NewTaobaopicturegetRequest() *TaobaopicturegetAPIRequest {
	return &TaobaopicturegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopicturegetAPIRequest) GetApiMethodName() string {
	return "taobao.picture.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopicturegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopicturegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleted is Deleted Setter
// 是否删除,unfroze代表没有删除
func (r *TaobaopicturegetAPIRequest) SetDeleted(_deleted string) error {
	r._deleted = _deleted
	r.Set("deleted", _deleted)
	return nil
}

// GetDeleted Deleted Getter
func (r TaobaopicturegetAPIRequest) GetDeleted() string {
	return r._deleted
}

// SetModifiedTime is ModifiedTime Setter
// 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
func (r *TaobaopicturegetAPIRequest) SetModifiedTime(_modifiedTime string) error {
	r._modifiedTime = _modifiedTime
	r.Set("modified_time", _modifiedTime)
	return nil
}

// GetModifiedTime ModifiedTime Getter
func (r TaobaopicturegetAPIRequest) GetModifiedTime() string {
	return r._modifiedTime
}

// SetTitle is Title Setter
// 图片标题,最大长度50字符,中英文都算一字符
func (r *TaobaopicturegetAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaopicturegetAPIRequest) GetTitle() string {
	return r._title
}

// SetStartDate is StartDate Setter
// 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaopicturegetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaopicturegetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaopicturegetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaopicturegetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetClientType is ClientType Setter
// 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
func (r *TaobaopicturegetAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r TaobaopicturegetAPIRequest) GetClientType() string {
	return r._clientType
}

// SetOrderBy is OrderBy Setter
// 图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc
func (r *TaobaopicturegetAPIRequest) SetOrderBy(_orderBy string) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r TaobaopicturegetAPIRequest) GetOrderBy() string {
	return r._orderBy
}

// SetUrls is Urls Setter
// 图片url查询接口
func (r *TaobaopicturegetAPIRequest) SetUrls(_urls string) error {
	r._urls = _urls
	r.Set("urls", _urls)
	return nil
}

// GetUrls Urls Getter
func (r TaobaopicturegetAPIRequest) GetUrls() string {
	return r._urls
}

// SetPictureId is PictureId Setter
// 图片ID
func (r *TaobaopicturegetAPIRequest) SetPictureId(_pictureId int64) error {
	r._pictureId = _pictureId
	r.Set("picture_id", _pictureId)
	return nil
}

// GetPictureId PictureId Getter
func (r TaobaopicturegetAPIRequest) GetPictureId() int64 {
	return r._pictureId
}

// SetPictureCategoryId is PictureCategoryId Setter
// 图片分类ID
func (r *TaobaopicturegetAPIRequest) SetPictureCategoryId(_pictureCategoryId int64) error {
	r._pictureCategoryId = _pictureCategoryId
	r.Set("picture_category_id", _pictureCategoryId)
	return nil
}

// GetPictureCategoryId PictureCategoryId Getter
func (r TaobaopicturegetAPIRequest) GetPictureCategoryId() int64 {
	return r._pictureCategoryId
}

// SetPageNo is PageNo Setter
// 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1
func (r *TaobaopicturegetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaopicturegetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数.每页返回最多返回100条,默认值40
func (r *TaobaopicturegetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaopicturegetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetIsHttps is IsHttps Setter
// 是否获取https的链接
func (r *TaobaopicturegetAPIRequest) SetIsHttps(_isHttps bool) error {
	r._isHttps = _isHttps
	r.Set("is_https", _isHttps)
	return nil
}

// GetIsHttps IsHttps Getter
func (r TaobaopicturegetAPIRequest) GetIsHttps() bool {
	return r._isHttps
}
