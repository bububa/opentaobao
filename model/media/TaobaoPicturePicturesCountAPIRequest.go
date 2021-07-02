package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPicturePicturesCountAPIRequest 图片总数查询 API请求
// taobao.picture.pictures.count
//
// 图片总数查询
type TaobaoPicturePicturesCountAPIRequest struct {
	model.Params
	// 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
	_startDate string
	// 图片分类
	_pictureCategoryId int64
	// 文件名
	_title string
	// 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
	_endDate string
	// 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
	_startModifiedDate string
	// 是否删除,undeleted代表没有删除,deleted表示删除
	_deleted string
	// 图片ID
	_pictureId int64
	// 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
	_clientType string
}

// NewTaobaoPicturePicturesCountRequest 初始化TaobaoPicturePicturesCountAPIRequest对象
func NewTaobaoPicturePicturesCountRequest() *TaobaoPicturePicturesCountAPIRequest {
	return &TaobaoPicturePicturesCountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPicturePicturesCountAPIRequest) GetApiMethodName() string {
	return "taobao.picture.pictures.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPicturePicturesCountAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStartDate is StartDate Setter
// 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaoPicturePicturesCountAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoPicturePicturesCountAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetPictureCategoryId is PictureCategoryId Setter
// 图片分类
func (r *TaobaoPicturePicturesCountAPIRequest) SetPictureCategoryId(_pictureCategoryId int64) error {
	r._pictureCategoryId = _pictureCategoryId
	r.Set("picture_category_id", _pictureCategoryId)
	return nil
}

// GetPictureCategoryId PictureCategoryId Getter
func (r TaobaoPicturePicturesCountAPIRequest) GetPictureCategoryId() int64 {
	return r._pictureCategoryId
}

// SetTitle is Title Setter
// 文件名
func (r *TaobaoPicturePicturesCountAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoPicturePicturesCountAPIRequest) GetTitle() string {
	return r._title
}

// SetEndDate is EndDate Setter
// 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaoPicturePicturesCountAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoPicturePicturesCountAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetStartModifiedDate is StartModifiedDate Setter
// 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
func (r *TaobaoPicturePicturesCountAPIRequest) SetStartModifiedDate(_startModifiedDate string) error {
	r._startModifiedDate = _startModifiedDate
	r.Set("start_modified_date", _startModifiedDate)
	return nil
}

// GetStartModifiedDate StartModifiedDate Getter
func (r TaobaoPicturePicturesCountAPIRequest) GetStartModifiedDate() string {
	return r._startModifiedDate
}

// SetDeleted is Deleted Setter
// 是否删除,undeleted代表没有删除,deleted表示删除
func (r *TaobaoPicturePicturesCountAPIRequest) SetDeleted(_deleted string) error {
	r._deleted = _deleted
	r.Set("deleted", _deleted)
	return nil
}

// GetDeleted Deleted Getter
func (r TaobaoPicturePicturesCountAPIRequest) GetDeleted() string {
	return r._deleted
}

// SetPictureId is PictureId Setter
// 图片ID
func (r *TaobaoPicturePicturesCountAPIRequest) SetPictureId(_pictureId int64) error {
	r._pictureId = _pictureId
	r.Set("picture_id", _pictureId)
	return nil
}

// GetPictureId PictureId Getter
func (r TaobaoPicturePicturesCountAPIRequest) GetPictureId() int64 {
	return r._pictureId
}

// SetClientType is ClientType Setter
// 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
func (r *TaobaoPicturePicturesCountAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r TaobaoPicturePicturesCountAPIRequest) GetClientType() string {
	return r._clientType
}
