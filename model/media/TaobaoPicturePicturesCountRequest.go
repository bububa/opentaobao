package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片总数查询 API请求
taobao.picture.pictures.count

图片总数查询
*/
type TaobaoPicturePicturesCountRequest struct {
    model.Params
    // 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
    _startDate   string
    // 图片分类
    _pictureCategoryId   int64
    // 文件名
    _title   string
    // 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
    _endDate   string
    // 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
    _startModifiedDate   string
    // 是否删除,undeleted代表没有删除,deleted表示删除
    _deleted   string
    // 图片ID
    _pictureId   int64
    // 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
    _clientType   string
}

// 初始化TaobaoPicturePicturesCountRequest对象
func NewTaobaoPicturePicturesCountRequest() *TaobaoPicturePicturesCountRequest{
    return &TaobaoPicturePicturesCountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPicturePicturesCountRequest) GetApiMethodName() string {
    return "taobao.picture.pictures.count"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPicturePicturesCountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaoPicturePicturesCountRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoPicturePicturesCountRequest) GetStartDate() string {
    return r._startDate
}
// PictureCategoryId Setter
// 图片分类
func (r *TaobaoPicturePicturesCountRequest) SetPictureCategoryId(_pictureCategoryId int64) error {
    r._pictureCategoryId = _pictureCategoryId
    r.Set("picture_category_id", _pictureCategoryId)
    return nil
}

// PictureCategoryId Getter
func (r TaobaoPicturePicturesCountRequest) GetPictureCategoryId() int64 {
    return r._pictureCategoryId
}
// Title Setter
// 文件名
func (r *TaobaoPicturePicturesCountRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TaobaoPicturePicturesCountRequest) GetTitle() string {
    return r._title
}
// EndDate Setter
// 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaoPicturePicturesCountRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoPicturePicturesCountRequest) GetEndDate() string {
    return r._endDate
}
// StartModifiedDate Setter
// 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
func (r *TaobaoPicturePicturesCountRequest) SetStartModifiedDate(_startModifiedDate string) error {
    r._startModifiedDate = _startModifiedDate
    r.Set("start_modified_date", _startModifiedDate)
    return nil
}

// StartModifiedDate Getter
func (r TaobaoPicturePicturesCountRequest) GetStartModifiedDate() string {
    return r._startModifiedDate
}
// Deleted Setter
// 是否删除,undeleted代表没有删除,deleted表示删除
func (r *TaobaoPicturePicturesCountRequest) SetDeleted(_deleted string) error {
    r._deleted = _deleted
    r.Set("deleted", _deleted)
    return nil
}

// Deleted Getter
func (r TaobaoPicturePicturesCountRequest) GetDeleted() string {
    return r._deleted
}
// PictureId Setter
// 图片ID
func (r *TaobaoPicturePicturesCountRequest) SetPictureId(_pictureId int64) error {
    r._pictureId = _pictureId
    r.Set("picture_id", _pictureId)
    return nil
}

// PictureId Getter
func (r TaobaoPicturePicturesCountRequest) GetPictureId() int64 {
    return r._pictureId
}
// ClientType Setter
// 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
func (r *TaobaoPicturePicturesCountRequest) SetClientType(_clientType string) error {
    r._clientType = _clientType
    r.Set("client_type", _clientType)
    return nil
}

// ClientType Getter
func (r TaobaoPicturePicturesCountRequest) GetClientType() string {
    return r._clientType
}
