package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPictureGetAPIRequest
获取图片信息 API请求
taobao.picture.get

获取图片信息 */
type TaobaoPictureGetAPIRequest struct {
	model.Params
	// 图片ID
	_pictureId int64
	// 图片分类ID
	_pictureCategoryId int64
	// 是否删除,unfroze代表没有删除
	_deleted string
	// 图片标题,最大长度50字符,中英文都算一字符
	_title string
	// 图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc
	_orderBy string
	// 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
	_startDate string
	// 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
	_endDate string
	// 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1
	_pageNo int64
	// 每页条数.每页返回最多返回100条,默认值40
	_pageSize int64
	// 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
	_modifiedTime string
	// 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
	_clientType string
	// 图片url查询接口
	_urls string
	// 是否获取https的链接
	_isHttps bool
}

// NewTaobaoPictureGetRequest 初始化TaobaoPictureGetAPIRequest对象
func NewTaobaoPictureGetRequest() *TaobaoPictureGetAPIRequest {
	return &TaobaoPictureGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPictureGetAPIRequest) GetApiMethodName() string {
	return "taobao.picture.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPictureGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PictureId Setter
// 图片ID
func (r *TaobaoPictureGetAPIRequest) SetPictureId(_pictureId int64) error {
	r._pictureId = _pictureId
	r.Set("picture_id", _pictureId)
	return nil
}

// Get PictureId Getter
func (r TaobaoPictureGetAPIRequest) GetPictureId() int64 {
	return r._pictureId
}

// Set is PictureCategoryId Setter
// 图片分类ID
func (r *TaobaoPictureGetAPIRequest) SetPictureCategoryId(_pictureCategoryId int64) error {
	r._pictureCategoryId = _pictureCategoryId
	r.Set("picture_category_id", _pictureCategoryId)
	return nil
}

// Get PictureCategoryId Getter
func (r TaobaoPictureGetAPIRequest) GetPictureCategoryId() int64 {
	return r._pictureCategoryId
}

// Set is Deleted Setter
// 是否删除,unfroze代表没有删除
func (r *TaobaoPictureGetAPIRequest) SetDeleted(_deleted string) error {
	r._deleted = _deleted
	r.Set("deleted", _deleted)
	return nil
}

// Get Deleted Getter
func (r TaobaoPictureGetAPIRequest) GetDeleted() string {
	return r._deleted
}

// Set is Title Setter
// 图片标题,最大长度50字符,中英文都算一字符
func (r *TaobaoPictureGetAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// Get Title Getter
func (r TaobaoPictureGetAPIRequest) GetTitle() string {
	return r._title
}

// Set is OrderBy Setter
// 图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc
func (r *TaobaoPictureGetAPIRequest) SetOrderBy(_orderBy string) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// Get OrderBy Getter
func (r TaobaoPictureGetAPIRequest) GetOrderBy() string {
	return r._orderBy
}

// Set is StartDate Setter
// 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaoPictureGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// Get StartDate Getter
func (r TaobaoPictureGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// Set is EndDate Setter
// 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaoPictureGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r TaobaoPictureGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is PageNo Setter
// 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1
func (r *TaobaoPictureGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoPictureGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 每页条数.每页返回最多返回100条,默认值40
func (r *TaobaoPictureGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoPictureGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is ModifiedTime Setter
// 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
func (r *TaobaoPictureGetAPIRequest) SetModifiedTime(_modifiedTime string) error {
	r._modifiedTime = _modifiedTime
	r.Set("modified_time", _modifiedTime)
	return nil
}

// Get ModifiedTime Getter
func (r TaobaoPictureGetAPIRequest) GetModifiedTime() string {
	return r._modifiedTime
}

// Set is ClientType Setter
// 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
func (r *TaobaoPictureGetAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// Get ClientType Getter
func (r TaobaoPictureGetAPIRequest) GetClientType() string {
	return r._clientType
}

// Set is Urls Setter
// 图片url查询接口
func (r *TaobaoPictureGetAPIRequest) SetUrls(_urls string) error {
	r._urls = _urls
	r.Set("urls", _urls)
	return nil
}

// Get Urls Getter
func (r TaobaoPictureGetAPIRequest) GetUrls() string {
	return r._urls
}

// Set is IsHttps Setter
// 是否获取https的链接
func (r *TaobaoPictureGetAPIRequest) SetIsHttps(_isHttps bool) error {
	r._isHttps = _isHttps
	r.Set("is_https", _isHttps)
	return nil
}

// Get IsHttps Getter
func (r TaobaoPictureGetAPIRequest) GetIsHttps() bool {
	return r._isHttps
}
