package main

import "github.com/json-iterator/go"

type GhostData struct {
	AppFields   []interface{} `json:"app_fields"`
	AppSettings []interface{} `json:"app_settings"`
	Apps        []interface{} `json:"apps"`
	Brute       []struct {
		Count        int    `json:"count"`
		FirstRequest int    `json:"firstRequest"`
		Key          string `json:"key"`
		LastRequest  int    `json:"lastRequest"`
		Lifetime     int    `json:"lifetime"`
	} `json:"brute"`
	Invites    []interface{} `json:"invites"`
	Migrations []struct {
		CurrentVersion string `json:"currentVersion"`
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Version        string `json:"version"`
	} `json:"migrations"`
	MigrationsLock []struct {
		AcquiredAt string `json:"acquired_at"`
		LockKey    string `json:"lock_key"`
		Locked     int    `json:"locked"`
		ReleasedAt string `json:"released_at"`
	} `json:"migrations_lock"`
	Permissions []struct {
		ActionType string      `json:"action_type"`
		CreatedAt  string      `json:"created_at"`
		CreatedBy  string      `json:"created_by"`
		ID         string      `json:"id"`
		Name       string      `json:"name"`
		ObjectID   interface{} `json:"object_id"`
		ObjectType string      `json:"object_type"`
		UpdatedAt  string      `json:"updated_at"`
		UpdatedBy  string      `json:"updated_by"`
	} `json:"permissions"`
	PermissionsApps  []interface{} `json:"permissions_apps"`
	PermissionsRoles []struct {
		ID           string `json:"id"`
		PermissionID string `json:"permission_id"`
		RoleID       string `json:"role_id"`
	} `json:"permissions_roles"`
	PermissionsUsers []interface{} `json:"permissions_users"`
	Posts            []struct {
		Amp                string      `json:"amp"`
		AuthorID           string      `json:"author_id"`
		CodeinjectionFoot  interface{} `json:"codeinjection_foot"`
		CodeinjectionHead  interface{} `json:"codeinjection_head"`
		CreatedAt          string      `json:"created_at"`
		CreatedBy          string      `json:"created_by"`
		CustomExcerpt      interface{} `json:"custom_excerpt"`
		CustomTemplate     interface{} `json:"custom_template"`
		FeatureImage       string      `json:"feature_image"`
		Featured           int         `json:"featured"`
		HTML               string      `json:"html"`
		ID                 string      `json:"id"`
		Locale             interface{} `json:"locale"`
		MetaDescription    interface{} `json:"meta_description"`
		MetaTitle          interface{} `json:"meta_title"`
		Mobiledoc          string      `json:"mobiledoc"`
		OgDescription      interface{} `json:"og_description"`
		OgImage            interface{} `json:"og_image"`
		OgTitle            interface{} `json:"og_title"`
		Page               int         `json:"page"`
		Plaintext          string      `json:"plaintext"`
		PublishedAt        string      `json:"published_at"`
		PublishedBy        string      `json:"published_by"`
		Slug               string      `json:"slug"`
		Status             string      `json:"status"`
		Title              string      `json:"title"`
		TwitterDescription interface{} `json:"twitter_description"`
		TwitterImage       interface{} `json:"twitter_image"`
		TwitterTitle       interface{} `json:"twitter_title"`
		UpdatedAt          string      `json:"updated_at"`
		UpdatedBy          string      `json:"updated_by"`
		UUID               string      `json:"uuid"`
		Visibility         string      `json:"visibility"`
	} `json:"posts"`
	PostsAuthors []struct {
		AuthorID  string `json:"author_id"`
		ID        string `json:"id"`
		PostID    string `json:"post_id"`
		SortOrder int    `json:"sort_order"`
	} `json:"posts_authors"`
	PostsTags []struct {
		ID        string `json:"id"`
		PostID    string `json:"post_id"`
		SortOrder int    `json:"sort_order"`
		TagID     string `json:"tag_id"`
	} `json:"posts_tags"`
	Roles []struct {
		CreatedAt   string `json:"created_at"`
		CreatedBy   string `json:"created_by"`
		Description string `json:"description"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		UpdatedAt   string `json:"updated_at"`
		UpdatedBy   string `json:"updated_by"`
	} `json:"roles"`
	RolesUsers []struct {
		ID     string `json:"id"`
		RoleID string `json:"role_id"`
		UserID string `json:"user_id"`
	} `json:"roles_users"`
	Settings []struct {
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
		ID        string `json:"id"`
		Key       string `json:"key"`
		Type      string `json:"type"`
		UpdatedAt string `json:"updated_at"`
		UpdatedBy string `json:"updated_by"`
		Value     string `json:"value"`
	} `json:"settings"`
	Subscribers []interface{} `json:"subscribers"`
	Tags        []struct {
		CreatedAt       string      `json:"created_at"`
		CreatedBy       string      `json:"created_by"`
		Description     interface{} `json:"description"`
		FeatureImage    interface{} `json:"feature_image"`
		ID              string      `json:"id"`
		MetaDescription interface{} `json:"meta_description"`
		MetaTitle       interface{} `json:"meta_title"`
		Name            string      `json:"name"`
		ParentID        interface{} `json:"parent_id"`
		Slug            string      `json:"slug"`
		UpdatedAt       string      `json:"updated_at"`
		UpdatedBy       string      `json:"updated_by"`
		Visibility      string      `json:"visibility"`
	} `json:"tags"`
	Users []struct {
		Accessibility        string      `json:"accessibility"`
		Bio                  interface{} `json:"bio"`
		CoverImage           interface{} `json:"cover_image"`
		CreatedAt            string      `json:"created_at"`
		CreatedBy            string      `json:"created_by"`
		Email                string      `json:"email"`
		Facebook             interface{} `json:"facebook"`
		GhostAuthAccessToken interface{} `json:"ghost_auth_access_token"`
		GhostAuthID          interface{} `json:"ghost_auth_id"`
		ID                   string      `json:"id"`
		LastSeen             string      `json:"last_seen"`
		Locale               interface{} `json:"locale"`
		Location             string      `json:"location"`
		MetaDescription      interface{} `json:"meta_description"`
		MetaTitle            interface{} `json:"meta_title"`
		Name                 string      `json:"name"`
		Password             string      `json:"password"`
		ProfileImage         string      `json:"profile_image"`
		Slug                 string      `json:"slug"`
		Status               string      `json:"status"`
		Tour                 string      `json:"tour"`
		Twitter              interface{} `json:"twitter"`
		UpdatedAt            string      `json:"updated_at"`
		UpdatedBy            string      `json:"updated_by"`
		Visibility           string      `json:"visibility"`
		Website              interface{} `json:"website"`
	} `json:"users"`
	Webhooks []interface{} `json:"webhooks"`
}

type GhostJSON struct {
	Db []struct {
		Data GhostData `json:"data"`
		Meta struct {
			ExportedOn int    `json:"exported_on"`
			Version    string `json:"version"`
		} `json:"meta"`
	} `json:"db"`
}

func (g GhostData) getPostTags(postID string) []string {
	var tags []string
	for _, v := range g.PostsTags {
		if v.PostID == postID {
			var tagName string
			for _, j := range g.Tags {
				if j.ID == v.TagID {
					tagName = j.Name
					break
				}
			}
			tags = append(tags, tagName)
		}
	}
	return tags
}

type MobileDoc struct {
	Atoms    []interface{}           `json:"atoms"`
	Cards    [][]jsoniter.RawMessage `json:"cards"`
	Markups  []interface{}           `json:"markups"`
	Sections [][]int                 `json:"sections"`
	Version  string                  `json:"version"`
}

type Card struct {
	CardName string `json:"cardName"`
	Markdown string `json:"markdown"`
}
