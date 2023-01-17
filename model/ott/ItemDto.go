package ott

// ItemDto 结构体
type ItemDto struct {
	// availablelanguages
	Availablelanguages []AvailableLanguageDto `json:"availablelanguages,omitempty" xml:"availablelanguages>available_language_dto,omitempty"`
	// availablesubtitles
	Availablesubtitles []AvailableSubtitleDto `json:"availablesubtitles,omitempty" xml:"availablesubtitles>available_subtitle_dto,omitempty"`
	// awardsandfestivals
	Awardsandfestivals []AwardSandFestivalDto `json:"awardsandfestivals,omitempty" xml:"awardsandfestivals>award_sand_festival_dto,omitempty"`
	// casts
	Casts []MemberDto `json:"casts,omitempty" xml:"casts>member_dto,omitempty"`
	// certifications
	Certifications []CertificationDto `json:"certifications,omitempty" xml:"certifications>certification_dto,omitempty"`
	// crews
	Crews []MemberDto `json:"crews,omitempty" xml:"crews>member_dto,omitempty"`
	// genres
	Genres []string `json:"genres,omitempty" xml:"genres>string,omitempty"`
	// images
	Images []ImageDto `json:"images,omitempty" xml:"images>image_dto,omitempty"`
	// ratings
	Ratings []RatingDto `json:"ratings,omitempty" xml:"ratings>rating_dto,omitempty"`
	// viewingoptions
	Viewingoptions []ViewingOptionDto `json:"viewingoptions,omitempty" xml:"viewingoptions>viewing_option_dto,omitempty"`
	// album
	Album string `json:"album,omitempty" xml:"album,omitempty"`
	// artist
	Artist string `json:"artist,omitempty" xml:"artist,omitempty"`
	// bitRate
	BitRate string `json:"bit_rate,omitempty" xml:"bit_rate,omitempty"`
	// description
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// duration
	Duration string `json:"duration,omitempty" xml:"duration,omitempty"`
	// procom
	Procom string `json:"procom,omitempty" xml:"procom,omitempty"`
	// programType
	ProgramType string `json:"program_type,omitempty" xml:"program_type,omitempty"`
	// resolution
	Resolution string `json:"resolution,omitempty" xml:"resolution,omitempty"`
	// seasonId
	SeasonId string `json:"season_id,omitempty" xml:"season_id,omitempty"`
	// seriesId
	SeriesId string `json:"series_id,omitempty" xml:"series_id,omitempty"`
	// subProgramType
	SubProgramType string `json:"sub_program_type,omitempty" xml:"sub_program_type,omitempty"`
	// subTitle
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// title
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// titleId
	TitleId string `json:"title_id,omitempty" xml:"title_id,omitempty"`
	// trailer
	Trailer string `json:"trailer,omitempty" xml:"trailer,omitempty"`
	// year
	Year string `json:"year,omitempty" xml:"year,omitempty"`
	// episodeNumber
	EpisodeNumber int64 `json:"episode_number,omitempty" xml:"episode_number,omitempty"`
	// seasonNumber
	SeasonNumber int64 `json:"season_number,omitempty" xml:"season_number,omitempty"`
}
