package models

type Issue struct {
	Total  int `json:"total"`
	Issues []struct {
		ID     string `json:"id"`
		Self   string `json:"self"`
		Key    string `json:"key"`
		Fields struct {
			Assignee struct {
				EmailAddress string `json:"emailAddress"`
			} `json:"assignee"`
			Subtasks []struct {
				Key    string `json:"key"`
				Fields struct {
					Summary string `json:"summary"`
					Status  struct {
						Description    string `json:"description"`
						Name           string `json:"name"`
						StatusCategory struct {
							Name string `json:"name"`
						} `json:"statusCategory"`
					} `json:"status"`
					Priority struct {
						Name string `json:"name"`
					} `json:"priority"`
					Issuetype struct {
						Description string `json:"description"`
						Name        string `json:"name"`
					} `json:"issuetype"`
				} `json:"fields"`
			} `json:"subtasks"`
			Summary string `json:"summary"`
			Creator struct {
				EmailAddress string `json:"emailAddress"`
			} `json:"creator"`
			Reporter struct {
				EmailAddress string `json:"emailAddress"`
			} `json:"reporter"`
			Progress struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
			} `json:"progress"`
			Status struct {
				Description    string `json:"description"`
				Name           string `json:"name"`
				StatusCategory struct {
					Name string `json:"name"`
				} `json:"statusCategory"`
			} `json:"status"`
			FixVersions []struct {
				Name string `json:"name"`
			} `json:"fixVersions"`
			Customfield10308 struct {
				EmailAddress string `json:"emailAddress"`
			} `json:"customfield_10308"`
			Customfield10307 struct {
				EmailAddress string `json:"emailAddress"`
			} `json:"customfield_10307"`
			Labels []string `json:"labels"`
		} `json:"fields"`
	}
}
