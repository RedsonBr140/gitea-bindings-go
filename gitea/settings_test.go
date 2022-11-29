package gitea

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_APISettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/settings/api", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprintf(w, `{"default_git_trees_per_page": 3,"max_response_items":1,"default_paging_num":2,"default_max_blob_size":1998}`)
	})

	APISettings, err := client.Settings.GetAPISettings()
	if err != nil {
		t.Errorf("Settings.GetAPISettings() returned error: %s", err)
	}

	want := &GeneralAPISettings{
		DefaultGitTreesPerPage: 3,
		MaxResponseItems:       1,
		DefaultPagingNum:       2,
		DefaultMaxBlobSize:     1998,
	}

	if !cmp.Equal(APISettings, want) {
		t.Errorf("Settings.GetAPISettings() returned %v, want %v", APISettings, want)
	}
}

func Test_AttachmentSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/settings/attachment", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprintf(w, `{"allowed_types":"*/*","enabled":true,"max_files":19,"max_size":2008}`)
	})

	Attachment, err := client.Settings.GetAttachmentSettings()
	if err != nil {
		t.Errorf("Settings.GetAttachmentSettings() returned error: %s", err)
	}

	want := &GeneralAttachmentSettings{
		AllowedTypes: "*/*",
		Enabled:      true,
		MaxFiles:     int64(19),
		MaxSize:      int64(2008),
	}

	if !cmp.Equal(Attachment, want) {
		t.Errorf("Settings.GetAttachmentSettings() returned %v, want %v", Attachment, want)
	}
}

func TestAPISettings_Marshal(t *testing.T) {
	testJSONMarshal(t, &GeneralAPISettings{}, "{}")

	u := &GeneralAPISettings{
		MaxResponseItems:       1,
		DefaultPagingNum:       2,
		DefaultGitTreesPerPage: 3,
		DefaultMaxBlobSize:     1998,
	}
	want := `{
	"max_response_items":1,
	"default_paging_num":2,
	"default_git_trees_per_page":3,
	"default_max_blob_size":1998
	}`

	testJSONMarshal(t, u, want)
}

func TestAttachmentSettings_Marshal(t *testing.T) {
	testJSONMarshal(t, &GeneralAttachmentSettings{}, "{}")

	u := &GeneralAttachmentSettings{
		Enabled:      true,
		AllowedTypes: "*/*",
		MaxSize:      100,
		MaxFiles:     20,
	}

	want := `{
	"enabled":true,
	"allowed_types":"*/*",
	"max_size":100,
	"max_files":20
	}`

	testJSONMarshal(t, u, want)
}

func TestRepoSettings_Marshal(t *testing.T) {
	testJSONMarshal(t, &GeneralRepoSettings{}, "{}")

	u := &GeneralRepoSettings{
		HttpGitDisabled:      true,
		LfsDisabled:          false,
		MigrationsDisabled:   true,
		StarsDisabled:        false,
		TimeTrackingDisabled: false,
	}

	want := `{
	"http_git_disabled":true,
	"lfs_disabled":false,
	"migrations_disabled":true,
	"stars_disabled":false,
	"time_tracking_disabled":false
	}`

	testJSONMarshal(t, u, want)
}

func TestUISettings_Marshal(t *testing.T) {
	testJSONMarshal(t, &GeneralUISettings{}, "{}")

	u := &GeneralUISettings{
		AllowedReactions: []string{"+1", "-1", "trollface"},
		CustomEmojis:     []string{"curitibatm"},
		DefaultTheme:     "blue",
	}

	want := `{
	"default_theme":"blue",
	"allowed_reactions":["+1","-1","trollface"],
	"custom_emojis":["curitibatm"]
	}`

	testJSONMarshal(t, u, want)

}
