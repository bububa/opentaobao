package pur

// MediaStatisticsDTO 
type MediaStatisticsDTO struct {
    // 文章摘要
    Summary   string `json:"summary,omitempty" xml:"summary,omitempty"`
    // 榜单排名
    Rowno   string `json:"rowno,omitempty" xml:"rowno,omitempty"`
    // 头条点赞数
    TopAvgLikenum   string `json:"top_avg_likenum,omitempty" xml:"top_avg_likenum,omitempty"`
    // 榜单所属月份
    ResultMonth   string `json:"result_month,omitempty" xml:"result_month,omitempty"`
    // 榜单排名变化
    RownoUp   string `json:"rowno_up,omitempty" xml:"rowno_up,omitempty"`
    // 平均评论数
    AvgCommentsNum   string `json:"avg_comments_num,omitempty" xml:"avg_comments_num,omitempty"`
    // 榜单指数变化
    CiUp   string `json:"ci_up,omitempty" xml:"ci_up,omitempty"`
    // 品类名称
    CategoryName   string `json:"category_name,omitempty" xml:"category_name,omitempty"`
    // 文章标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 阅读数
    ReadNum   string `json:"read_num,omitempty" xml:"read_num,omitempty"`
    // 昵称
    NickName   string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
    // 平均转发数
    AvgRepostsNum   string `json:"avg_reposts_num,omitempty" xml:"avg_reposts_num,omitempty"`
    // 名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 次条点赞数
    SecondAvgLikenum   string `json:"second_avg_likenum,omitempty" xml:"second_avg_likenum,omitempty"`
    // 原创比例
    CopyrightProportion   string `json:"copyright_proportion,omitempty" xml:"copyright_proportion,omitempty"`
    // 发布时间
    PostTime   string `json:"post_time,omitempty" xml:"post_time,omitempty"`
    // 头条阅读数
    TopAvgReadnum   string `json:"top_avg_readnum,omitempty" xml:"top_avg_readnum,omitempty"`
    // 转发数
    RepostsNum   string `json:"reposts_num,omitempty" xml:"reposts_num,omitempty"`
    // 发布次数
    UrlTimes   string `json:"url_times,omitempty" xml:"url_times,omitempty"`
    // 品类ID
    CategoryId   string `json:"category_id,omitempty" xml:"category_id,omitempty"`
    // 评论数
    CommentsNum   string `json:"comments_num,omitempty" xml:"comments_num,omitempty"`
    // 文章地址
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    // 文章正文
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    // 3-N条阅读数
    OtherAvgReadnum   string `json:"other_avg_readnum,omitempty" xml:"other_avg_readnum,omitempty"`
    // 超过10w的文章数
    UrlNumMoreTenw   string `json:"url_num_more_tenw,omitempty" xml:"url_num_more_tenw,omitempty"`
    // 好看数
    LikeNum   string `json:"like_num,omitempty" xml:"like_num,omitempty"`
    // 次条阅读数
    SecondAvgReadnum   string `json:"second_avg_readnum,omitempty" xml:"second_avg_readnum,omitempty"`
    // 榜单指数
    Ci   string `json:"ci,omitempty" xml:"ci,omitempty"`
    // 平均点赞数
    AvgLikeNum   string `json:"avg_like_num,omitempty" xml:"avg_like_num,omitempty"`
    // 原创发布数
    CopyrightUrlNum   string `json:"copyright_url_num,omitempty" xml:"copyright_url_num,omitempty"`
    // 发布文章数
    UrlNum   string `json:"url_num,omitempty" xml:"url_num,omitempty"`
    // 榜单分类排名
    TypesRowno   string `json:"types_rowno,omitempty" xml:"types_rowno,omitempty"`
    // 平均阅读数
    AvgReadNum   string `json:"avg_read_num,omitempty" xml:"avg_read_num,omitempty"`
    // 文章类型
    ArticleType   string `json:"article_type,omitempty" xml:"article_type,omitempty"`
    // 榜单类型
    BoardType   string `json:"board_type,omitempty" xml:"board_type,omitempty"`
    // 统计时间
    StatisticTime   string `json:"statistic_time,omitempty" xml:"statistic_time,omitempty"`
    // 发布类型
    PostType   string `json:"post_type,omitempty" xml:"post_type,omitempty"`
}
