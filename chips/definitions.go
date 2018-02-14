package chips

import "time"

type CompoResponse struct {
	Compo struct {
		ID           int         `json:"id"`
		Name         string      `json:"name"`
		PrimaryImage string      `json:"primary_image"`
		Description  string      `json:"description"`
		Theme        string      `json:"theme"`
		ThemeShort   string      `json:"theme_short"`
		Start        time.Time   `json:"start"`
		Stop         time.Time   `json:"stop"`
		JudgingStop  time.Time   `json:"judging_stop"`
		Released     interface{} `json:"released"`
		Entries      int         `json:"entries"`
		HeaderImage  string      `json:"header_image"`
		Meme         bool        `json:"meme"`
		Nogold       interface{} `json:"nogold"`
		Type         string      `json:"type"`
		Types        []string    `json:"types"`
		AllowVoting  bool        `json:"allow_voting"`
		KeyNames     []string    `json:"key_names"`
		StartDate    time.Time   `json:"startDate"`
		StopDate     time.Time   `json:"stopDate"`
		State        string      `json:"state"`
	} `json:"compo"`
	Images            []Image       `json:"images"`
	Entries           []Entry       `json:"entries"`
	FavoriteEntryIDs  []interface{} `json:"favoriteEntryIDs"`
	CommentedEntryIDs []interface{} `json:"commentedEntryIDs"`
	EntriesSubmitted  struct {
		Song  int `json:"song"`
		Art   int `json:"art"`
		Story int `json:"story"`
	} `json:"entriesSubmitted"`
	Categories []Category `json:"categories"`
	CanJudge   bool       `json:"can_judge"`
}

type Image struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	CompoID string `json:"compo_id"`
}

type Entry struct {
	ID            int         `json:"id"`
	UserID        int         `json:"user_id"`
	CompoID       int         `json:"compo_id"`
	Created       time.Time   `json:"created"`
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	URL           interface{} `json:"url"`
	Anonymous     bool        `json:"anonymous"`
	UploadedURL   string      `json:"uploaded_url"`
	Hidden        bool        `json:"hidden"`
	UserID1       int         `json:"user_id_1"`
	UserID2       int         `json:"user_id_2"`
	UserID3       int         `json:"user_id_3"`
	UserID4       int         `json:"user_id_4"`
	UserID5       int         `json:"user_id_5"`
	UserID0Conf   string      `json:"user_id_0_conf"`
	UserID1Conf   string      `json:"user_id_1_conf"`
	UserID2Conf   string      `json:"user_id_2_conf"`
	UserID3Conf   string      `json:"user_id_3_conf"`
	UserID4Conf   string      `json:"user_id_4_conf"`
	UserID5Conf   string      `json:"user_id_5_conf"`
	OverallScore  interface{} `json:"overall_score"`
	Plays         int         `json:"plays"`
	Downloads     int         `json:"downloads"`
	Downloadable  bool        `json:"downloadable"`
	Hearts        int         `json:"hearts"`
	Comments      int         `json:"comments"`
	Type          string      `json:"type"`
	Reviews       int         `json:"reviews"`
	FollowedTheme bool        `json:"followed_theme"`
	IsJudged      bool        `json:"is_judged"`
	ModComments   string      `json:"mod_comments"`
	IsJoke        bool        `json:"is_joke"`
	KeyNames      []string    `json:"key_names"`
}

type Category struct {
	ID        int    `json:"id"`
	CompoID   int    `json:"compo_id"`
	Category1 string `json:"category_1"`
	Category2 string `json:"category_2"`
	Category3 string `json:"category_3"`
	Category4 string `json:"category_4"`
	Category5 string `json:"category_5"`
	Category6 string `json:"category_6"`
	Category7 string `json:"category_7"`
	Category8 string `json:"category_8"`
	Category9 string `json:"category_9"`
	Type      string `json:"type"`
}
