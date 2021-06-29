package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片获取 API请求
taobao.picture.pictures.get

图片空间对外的图片获取接口，该接口只针对分页获取，获取某一页的图片，该接口不支持总数的查询asd
*/
type TaobaoPicturePicturesGetRequest struct {
    model.Params
    // 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
    startDate   string
    // 图片分类
    pictureCategoryId   int64
    // 图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc
    orderBy   string
    // 图片标题,最大长度50字符,中英文都算一字符
    title   string
    // 每页条数.每页返回最多返回100条,默认值40
    pageSize   int64
    // 结束时间
    endDate   string
    // 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1
    currentPage   int64
    // 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
    startModifiedDate   string
    // 是否删除,deleted代表没有删除
    deleted   string
    // 图片ID
    pictureId   int64
    // 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
    clientType   string
    // 图片url查询接口
    urls   string
    // 是否获取https的链接
    isHttps   bool
}

// 初始化TaobaoPicturePicturesGetRequest对象
func NewTaobaoPicturePicturesGetRequest() *TaobaoPicturePicturesGetRequest{
    return &TaobaoPicturePicturesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPicturePicturesGetRequest) GetApiMethodName() string {
    return "taobao.picture.pictures.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPicturePicturesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaoPicturePicturesGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoPicturePicturesGetRequest) GetStartDate() string {
    return r.startDate
}
// PictureCategoryId Setter
// 图片分类
func (r *TaobaoPicturePicturesGetRequest) SetPictureCategoryId(pictureCategoryId int64) error {
    r.pictureCategoryId = pictureCategoryId
    r.Set("picture_category_id", pictureCategoryId)
    return nil
}

// PictureCategoryId Getter
func (r TaobaoPicturePicturesGetRequest) GetPictureCategoryId() int64 {
    return r.pictureCategoryId
}
// OrderBy Setter
// 图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc
func (r *TaobaoPicturePicturesGetRequest) SetOrderBy(orderBy string) error {
    r.orderBy = orderBy
    r.Set("order_by", orderBy)
    return nil
}

// OrderBy Getter
func (r TaobaoPicturePicturesGetRequest) GetOrderBy() string {
    return r.orderBy
}
// Title Setter
// 图片标题,最大长度50字符,中英文都算一字符
func (r *TaobaoPicturePicturesGetRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoPicturePicturesGetRequest) GetTitle() string {
    return r.title
}
// PageSize Setter
// 每页条数.每页返回最多返回100条,默认值40
func (r *TaobaoPicturePicturesGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoPicturePicturesGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// EndDate Setter
// 结束时间
func (r *TaobaoPicturePicturesGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoPicturePicturesGetRequest) GetEndDate() string {
    return r.endDate
}
// CurrentPage Setter
// 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1
func (r *TaobaoPicturePicturesGetRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoPicturePicturesGetRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// StartModifiedDate Setter
// 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
func (r *TaobaoPicturePicturesGetRequest) SetStartModifiedDate(startModifiedDate string) error {
    r.startModifiedDate = startModifiedDate
    r.Set("start_modified_date", startModifiedDate)
    return nil
}

// StartModifiedDate Getter
func (r TaobaoPicturePicturesGetRequest) GetStartModifiedDate() string {
    return r.startModifiedDate
}
// Deleted Setter
// 是否删除,deleted代表没有删除
func (r *TaobaoPicturePicturesGetRequest) SetDeleted(deleted string) error {
    r.deleted = deleted
    r.Set("deleted", deleted)
    return nil
}

// Deleted Getter
func (r TaobaoPicturePicturesGetRequest) GetDeleted() string {
    return r.deleted
}
// PictureId Setter
// 图片ID
func (r *TaobaoPicturePicturesGetRequest) SetPictureId(pictureId int64) error {
    r.pictureId = pictureId
    r.Set("picture_id", pictureId)
    return nil
}

// PictureId Getter
func (r TaobaoPicturePicturesGetRequest) GetPictureId() int64 {
    return r.pictureId
}
// ClientType Setter
// 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
func (r *TaobaoPicturePicturesGetRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

// ClientType Getter
func (r TaobaoPicturePicturesGetRequest) GetClientType() string {
    return r.clientType
}
// Urls Setter
// 图片url查询接口
func (r *TaobaoPicturePicturesGetRequest) SetUrls(urls string) error {
    r.urls = urls
    r.Set("urls", urls)
    return nil
}

// Urls Getter
func (r TaobaoPicturePicturesGetRequest) GetUrls() string {
    return r.urls
}
// IsHttps Setter
// 是否获取https的链接
func (r *TaobaoPicturePicturesGetRequest) SetIsHttps(isHttps bool) error {
    r.isHttps = isHttps
    r.Set("is_https", isHttps)
    return nil
}

// IsHttps Getter
func (r TaobaoPicturePicturesGetRequest) GetIsHttps() bool {
    return r.isHttps
}
