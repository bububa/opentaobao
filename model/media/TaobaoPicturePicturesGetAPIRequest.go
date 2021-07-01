package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPicturePicturesGetAPIRequest
图片获取 API请求
taobao.picture.pictures.get

图片空间对外的图片获取接口，该接口只针对分页获取，获取某一页的图片，该接口不支持总数的查询asd */
type TaobaoPicturePicturesGetAPIRequest struct {
	model.Params
	// 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
	_startDate string
	// 图片分类
	_pictureCategoryId int64
	// 图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc
	_orderBy string
	// 图片标题,最大长度50字符,中英文都算一字符
	_title string
	// 每页条数.每页返回最多返回100条,默认值40
	_pageSize int64
	// 结束时间
	_endDate string
	// 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1
	_currentPage int64
	// 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
	_startModifiedDate string
	// 是否删除,deleted代表没有删除
	_deleted string
	// 图片ID
	_pictureId int64
	// 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
	_clientType string
	// 图片url查询接口
	_urls string
	// 是否获取https的链接
	_isHttps bool
}

// New
