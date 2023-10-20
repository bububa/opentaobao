package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopicturepicturescountAPIRequest 图片总数查询 API请求
// taobao.picture.pictures.count
//
// 图片总数查询，目前出于对数据库的保护暂不支持此功能
type TaobaopicturepicturescountAPIRequest struct {
	model.Params
	// 是否删除,undeleted代表没有删除,deleted表示删除
	_deleted string
	// 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
	_startModifiedDate string
	// 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
	_startDate string
	// 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
	_endDate string
	// 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
	_clientType string
	// 文件名
	_title string
	// 图片ID
	_pictureId int64
	// 图片分类
	_pictureCategoryId int64
}

// NewTaobaopicturepicturescountRequest 初始化TaobaopicturepicturescountAPIRequest对象
func NewTaobaopicturepicturescountRequest() *TaobaopicturepicturescountAPIRequest {
	return &TaobaopicturepicturescountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopicturepicturescountAPIRequest) GetApiMethodName() string {
	return "taobao.picture.pictures.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopicturepicturescountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopicturepicturescountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleted is Deleted Setter
// 是否删除,undeleted代表没有删除,deleted表示删除
func (r *TaobaopicturepicturescountAPIRequest) SetDeleted(_deleted string) error {
	r._deleted = _deleted
	r.Set("deleted", _deleted)
	return nil
}

// GetDeleted Deleted Getter
func (r TaobaopicturepicturescountAPIRequest) GetDeleted() string {
	return r._deleted
}

// SetStartModifiedDate is StartModifiedDate Setter
// 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
func (r *TaobaopicturepicturescountAPIRequest) SetStartModifiedDate(_startModifiedDate string) error {
	r._startModifiedDate = _startModifiedDate
	r.Set("start_modified_date", _startModifiedDate)
	return nil
}

// GetStartModifiedDate StartModifiedDate Getter
func (r TaobaopicturepicturescountAPIRequest) GetStartModifiedDate() string {
	return r._startModifiedDate
}

// SetStartDate is StartDate Setter
// 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaopicturepicturescountAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaopicturepicturescountAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaopicturepicturescountAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaopicturepicturescountAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetClientType is ClientType Setter
// 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
func (r *TaobaopicturepicturescountAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r TaobaopicturepicturescountAPIRequest) GetClientType() string {
	return r._clientType
}

// SetTitle is Title Setter
// 文件名
func (r *TaobaopicturepicturescountAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaopicturepicturescountAPIRequest) GetTitle() string {
	return r._title
}

// SetPictureId is PictureId Setter
// 图片ID
func (r *TaobaopicturepicturescountAPIRequest) SetPictureId(_pictureId int64) error {
	r._pictureId = _pictureId
	r.Set("picture_id", _pictureId)
	return nil
}

// GetPictureId PictureId Getter
func (r TaobaopicturepicturescountAPIRequest) GetPictureId() int64 {
	return r._pictureId
}

// SetPictureCategoryId is PictureCategoryId Setter
// 图片分类
func (r *TaobaopicturepicturescountAPIRequest) SetPictureCategoryId(_pictureCategoryId int64) error {
	r._pictureCategoryId = _pictureCategoryId
	r.Set("picture_category_id", _pictureCategoryId)
	return nil
}

// GetPictureCategoryId PictureCategoryId Getter
func (r TaobaopicturepicturescountAPIRequest) GetPictureCategoryId() int64 {
	return r._pictureCategoryId
}
