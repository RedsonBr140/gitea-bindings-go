# TODO List
"We've  got work to do"

# Bind API Endpoints
List of API Endpoints(v1.17.1+69-gb69472d) and the status

<details markdown="1">
<summary><input disabled type="checkbox">Admin</input></summary>

 - **GET**
   - [ ] /admin/cron
   - [ ] /admin/orgs
   - [ ] /admin/unadopted
   - [ ] /admin/users
 - **POST**
   - [ ] /admin/cron/{task}
   - [ ] /admin/unadopted/{owner}/{repo}
   - [ ] /admin/users
   - [ ] /admin/users/{username}/keys
   - [ ] /admin/users/{username}/orgs
   - [ ] /admin/users/{username}/repos
 - **DELETE**
   - [ ] /admin/unadopted/{owner}/{repo}
   - [ ] /admin/users/{username}
   - [ ] /admin/users/{username}/keys/{id}

</details>

<details>
<summary><input disabled type="checkbox">Miscellaneous</input></summary>

 - **GET**
   - [ ] /nodeinfo
   - [ ] /signing-key.gpg
   - [X] /version



 - **POST**
   - [X] /markdown
   - [X] /markdown/raw

</details>

<details>
<summary><input disabled type="checkbox">Notification</input></summary>

 - **GET**
   - [ ] /notifications
   - [ ] /notifications/new
   - [ ] /notifications/threads/{id}
   - [ ] /repos/{owner}/{repo}/notifications



 - **PUT**
   - [ ] /notifications
	 - [ ] /repos/{owner}/{repo}/notifications

 - **PATCH**
   - [ ] /notifications/threads/{id}
</details>

<details>
<summary><input disabled type="checkbox">Organization</input></summary>

 - **GET**
   - [ ] /orgs
   - [ ] /orgs/{org}
   - [ ] /orgs/{org}/hooks
   - [ ] /orgs/{org}/hooks/{id}
   - [ ] /orgs/{org}/labels
   - [ ] /orgs/{org}/labels/{id}
   - [ ] /orgs/{org}/members
   - [ ] /orgs/{org}/members/{username}
   - [ ] /orgs/{org}/public_members
   - [ ] /orgs/{org}/public_members/{username}
   - [ ] /orgs/{org}/repos
   - [ ] /orgs/{org}/teams
   - [ ] /orgs/{org}/teams/search
   - [ ] /teams/{id}
   - [ ] /teams/{id}/members
   - [ ] /teams/{id}/members/{username}
   - [ ] /teams/{id}/repos
   - [ ] /teams/{id}/repos/{org}/{repo}
   - [ ] /user/orgs
   - [ ] /users/{username}/orgs
   - [ ] /users/{username}/orgs/{org}/permissions
 - **POST**
   - [ ] /orgs
   - [ ] /orgs/{org}/hooks/
   - [ ] /orgs/{org}/labels
   - [ ] /orgs/{org}/repos
   - [ ] /orgs/{org}/teams
 - **PUT**
   - [ ] /orgs/{org}/public_members/{username}
   - [ ] /teams/{id}/members/{username}
   - [ ] /teams/{id}/repos/{org}/{repo}

 - **PATCH**
   - [ ] /teams/{id}
   - [ ] /orgs/{org}
   - [ ] /orgs/{org}/hooks/{id}
   - [ ] /orgs/{org}/labels/{id}
 - **DELETE**
   - [ ] /orgs/{org}
   - [ ] /orgs/{org}/hooks/{id}
   - [ ] /orgs/{org}/labels/{id}
   - [ ] /orgs/{org}/members/{username}
   - [ ] /orgs/{org}/public_members/{username}
   - [ ] /teams/{id}
   - [ ] /teams/{id}/members/{username}
   - [ ] /teams/{id}/repos/{org}/{repo}
</details>

<details>
<summary><input disabled type="checkbox">Package</input></summary>
 - **GET**
   - [ ] /packages/{owner}
   - [ ] /packages/{owner}/{type}/{name}/{version}
   - [ ] /packages/{owner}/{type}/{name}/{version}/files
 - **DELETE**
   - [ ] /packages/{owner}/{type}/{name}/{version}
</details>

<details>
<summary><input disabled type="checkbox">Issue</input></summary>

 - **GET**
   - [ ] /repos/issues/search
   - [ ] /repos/{owner}/{repo}/issues
   - [ ] /repos/{owner}/{repo}/issues/comments
   - [ ] /repos/{owner}/{repo}/issues/comments/{id}
   - [ ] /repos/{owner}/{repo}/issues/comments/{id}/reactions
   - [ ] /repos/{owner}/{repo}/issues/{index}
   - [ ] /repos/{owner}/{repo}/issues/{index}/comments
   - [ ] /repos/{owner}/{repo}/issues/{index}/labels
   - [ ] /repos/{owner}/{repo}/issues/{index}/reactions
   - [ ] /repos/{owner}/{repo}/issues/{index}/subscriptions
   - [ ] /repos/{owner}/{repo}/issues/{index}/subscriptions/check
   - [ ] /repos/{owner}/{repo}/issues/{index}/timeline
   - [ ] /repos/{owner}/{repo}/issues/{index}/times
   - [ ] /repos/{owner}/{repo}/labels
   - [ ] /repos/{owner}/{repo}/labels/{id}
   - [ ] /repos/{owner}/{repo}/milestones
   - [ ] /repos/{owner}/{repo}/milestones/{id}
 - **POST**
   - [ ] /repos/{owner}/{repo}/issues
   - [ ] /repos/{owner}/{repo}/issues/{index}/comments
   - [ ] /repos/{owner}/{repo}/issues/{index}/deadline
   - [ ] /repos/{owner}/{repo}/issues/{index}/labels
   - [ ] /repos/{owner}/{repo}/issues/{index}/reactions
   - [ ] /repos/{owner}/{repo}/issues/{index}/stopwatch/start
   - [ ] /repos/{owner}/{repo}/issues/{index}/stopwatch/stop
   - [ ] /repos/{owner}/{repo}/issues/{index}/times
   - [ ] /repos/{owner}/{repo}/labels
   - [ ] /repos/{owner}/{repo}/milestones
 - **PUT**
   - [ ] /repos/{owner}/{repo}/issues/{index}/labels
	 - [ ] /repos/{owner}/{repo}/issues/{index}/subscriptions/{user}
 - **PATCH**
   - [ ] /repos/{owner}/{repo}/milestones/{id}
   - [ ] /repos/{owner}/{repo}/labels/{id}
   - [ ] /repos/{owner}/{repo}/issues/{index}/comments/{id}
   - [ ] /repos/{owner}/{repo}/issues/{index}
   - [ ] /repos/{owner}/{repo}/issues/comments/{id}
 - **DELETE**
   - [ ] /repos/{owner}/{repo}/issues/comments/{id}
   - [ ] /repos/{owner}/{repo}/issues/comments/{id}/reactions
   - [ ] /repos/{owner}/{repo}/issues/{index}
   - [ ] /repos/{owner}/{repo}/issues/{index}/comments/{id}
   - [ ] /repos/{owner}/{repo}/issues/{index}/labels
   - [ ] /repos/{owner}/{repo}/issues/{index}/labels/{id}
   - [ ] /repos/{owner}/{repo}/issues/{index}/reactions
   - [ ] /repos/{owner}/{repo}/issues/{index}/stopwatch/delete
   - [ ] /repos/{owner}/{repo}/issues/{index}/subscriptions/{user}
   - [ ] /repos/{owner}/{repo}/issues/{index}/times
   - [ ] /repos/{owner}/{repo}/issues/{index}/times/{id}
   - [ ] /repos/{owner}/{repo}/labels/{id}
   - [ ] /repos/{owner}/{repo}/milestones/{id}
</details>

<details>
<summary><input disabled type="checkbox">Repository</input></summary>

- **GET**
   - [X] /repos/search
   - [X] /repos/{owner}/{repo}
   - [ ] /repos/{owner}/{repo}/archive/{archive}
   - [ ] /repos/{owner}/{repo}/assignees
   - [ ] /repos/{owner}/{repo}/branch_protections
   - [ ] /repos/{owner}/{repo}/branch_protections/{name}
   - [ ] /repos/{owner}/{repo}/branches
   - [ ] /repos/{owner}/{repo}/branches/{branch}
   - [ ] /repos/{owner}/{repo}/collaborators
   - [ ] /repos/{owner}/{repo}/collaborators/{collaborator}
   - [ ] /repos/{owner}/{repo}/collaborators/{collaborator}/permission
   - [ ] /repos/{owner}/{repo}/commits
   - [ ] /repos/{owner}/{repo}/commits/{ref}/status
   - [ ] /repos/{owner}/{repo}/commits/{ref}/statuses
   - [ ] /repos/{owner}/{repo}/contents
   - [ ] /repos/{owner}/{repo}/contents/{filepath}
   - [ ] /repos/{owner}/{repo}/editorconfig/{filepath}
   - [ ] /repos/{owner}/{repo}/forks
   - [ ] /repos/{owner}/{repo}/git/blobs/{sha}
   - [ ] /repos/{owner}/{repo}/git/commits/{sha}
   - [ ] /repos/{owner}/{repo}/git/commits/{sha}.{diffType}
   - [ ] /repos/{owner}/{repo}/git/notes/{sha}
   - [ ] /repos/{owner}/{repo}/git/refs
   - [ ] /repos/{owner}/{repo}/git/refs/{ref}
   - [ ] /repos/{owner}/{repo}/git/tags/{sha}
   - [ ] /repos/{owner}/{repo}/git/trees/{sha}
   - [ ] /repos/{owner}/{repo}/hooks
   - [ ] /repos/{owner}/{repo}/hooks/git
   - [ ] /repos/{owner}/{repo}/hooks/git/{id}
   - [ ] /repos/{owner}/{repo}/hooks/{id}
   - [ ] /repos/{owner}/{repo}/issue_templates
   - [ ] /repos/{owner}/{repo}/keys
   - [ ] /repos/{owner}/{repo}/keys/{id}
   - [ ] /repos/{owner}/{repo}/languages
   - [ ] /repos/{owner}/{repo}/media/{filepath}
   - [ ] /repos/{owner}/{repo}/pulls
   - [ ] /repos/{owner}/{repo}/pulls/{index}
   - [ ] /repos/{owner}/{repo}/pulls/{index}.{diffType}
   - [ ] /repos/{owner}/{repo}/pulls/{index}/commits
   - [ ] /repos/{owner}/{repo}/pulls/{index}/merge
   - [ ] /repos/{owner}/{repo}/pulls/{index}/reviews
   - [ ] /repos/{owner}/{repo}/pulls/{index}/reviews/{id}
   - [ ] /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/comments
   - [ ] /repos/{owner}/{repo}/raw/{filepath}
   - [ ] /repos/{owner}/{repo}/releases
   - [ ] /repos/{owner}/{repo}/releases/tags/{tag}
   - [ ] /repos/{owner}/{repo}/releases/{id}
   - [ ] /repos/{owner}/{repo}/releases/{id}/assets
   - [ ] /repos/{owner}/{repo}/releases/{id}/assets/{attachment_id}
   - [ ] /repos/{owner}/{repo}/reviewers
   - [ ] /repos/{owner}/{repo}/signing-key.gpg
   - [ ] /repos/{owner}/{repo}/stargazers
   - [ ] /repos/{owner}/{repo}/statuses/{sha}
   - [ ] /repos/{owner}/{repo}/subscribers
   - [ ] /repos/{owner}/{repo}/subscription
   - [ ] /repos/{owner}/{repo}/tags
   - [ ] /repos/{owner}/{repo}/tags/{tag}
   - [ ] /repos/{owner}/{repo}/teams
   - [ ] /repos/{owner}/{repo}/teams/{team}
   - [ ] /repos/{owner}/{repo}/times
   - [ ] /repos/{owner}/{repo}/times/{user}
   - [ ] /repos/{owner}/{repo}/topics
   - [ ] /repos/{owner}/{repo}/wiki/page/{pageName}
   - [ ] /repos/{owner}/{repo}/wiki/pages
   - [ ] /repos/{owner}/{repo}/wiki/revisions/{pageName}
   - [ ] /repositories/{id}
   - [ ] /topics/search
 - **POST**
   - [ ] /repos/migrate
   - [ ] /repos/{owner}/{repo}/branch_protections
   - [ ] /repos/{owner}/{repo}/branches
   - [ ] /repos/{owner}/{repo}/contents/{filepath}
   - [ ] /repos/{owner}/{repo}/diffpatch
   - [ ] /repos/{owner}/{repo}/forks
   - [ ] /repos/{owner}/{repo}/hooks
   - [ ] /repos/{owner}/{repo}/hooks/{id}/tests
   - [ ] /repos/{owner}/{repo}/keys
   - [ ] /repos/{owner}/{repo}/mirror-sync
   - [ ] /repos/{owner}/{repo}/pulls
   - [ ] /repos/{owner}/{repo}/pulls/{index}/merge
   - [ ] /repos/{owner}/{repo}/pulls/{index}/requested_reviewers
   - [ ] /repos/{owner}/{repo}/pulls/{index}/reviews
   - [ ] /repos/{owner}/{repo}/pulls/{index}/reviews/{id}
   - [ ] /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/dismissals
   - [ ] /repos/{owner}/{repo}/pulls/{index}/reviews/{id}/undismissals
   - [ ] /repos/{owner}/{repo}/pulls/{index}/update
   - [ ] /repos/{owner}/{repo}/releases
   - [ ] /repos/{owner}/{repo}/releases/{id}/assets
   - [ ] /repos/{owner}/{repo}/statuses/{sha}
   - [ ] /repos/{owner}/{repo}/tags
   - [ ] /repos/{owner}/{repo}/transfer
   - [ ] /repos/{owner}/{repo}/transfer/accept
   - [ ] /repos/{owner}/{repo}/transfer/reject
   - [ ] /repos/{owner}/{repo}/wiki/new
   - [ ] /repos/{template_owner}/{template_repo}/generate
 - **PUT**
   - [ ] /repos/{owner}/{repo}/collaborators/{collaborator}
   - [ ] /repos/{owner}/{repo}/contents/{filepath}
   - [ ] /repos/{owner}/{repo}/subscription
   - [ ] /repos/{owner}/{repo}/teams/{team}
   - [ ] /repos/{owner}/{repo}/topics
   - [ ] /repos/{owner}/{repo}/topics/{topic}
 - **PATCH**
   - [ ] /repos/{owner}/{repo}
   - [ ] /repos/{owner}/{repo}/branch_protections/{name}
   - [ ] /repos/{owner}/{repo}/hooks/git/{id}
   - [ ] /repos/{owner}/{repo}/hooks/{id}
   - [ ] /repos/{owner}/{repo}/pulls/{index}
   - [ ] /repos/{owner}/{repo}/releases/{id}
   - [ ] /repos/{owner}/{repo}/releases/{id}/assets/{attachment_id}
   - [ ] /repos/{owner}/{repo}/wiki/page/{pageName}
 - **DELETE**
   - [X] repos/{owner}/{repo}
   - [ ] repos/{owner}/{repo}/branch_protections/{name}
   - [ ] repos/{owner}/{repo}/collaborators/{collaborator}
   - [ ] repos/{owner}/{repo}/contents/{filepath}
   - [ ] repos/{owner}/{repo}/hooks/git/{id}
   - [ ] repos/{owner}/{repo}/hooks/{id}
   - [ ] repos/{owner}/{repo}/keys/{id}
   - [ ] repos/{owner}/{repo}/pulls/{index}/merge
   - [ ] repos/{owner}/{repo}/pulls/{index}/requested_reviewers
   - [ ] repos/{owner}/{repo}/pulls/{index}/reviews/{id}
   - [ ] repos/{owner}/{repo}/releases/tags/{tag}
   - [ ] repos/{owner}/{repo}/releases/{id}
   - [ ] repos/{owner}/{repo}/releases/{id}/assets/{attachment_id}
   - [ ] repos/{owner}/{repo}/subscription
   - [ ] repos/{owner}/{repo}/tags/{tag}
   - [ ] repos/{owner}/{repo}/teams/{team}
   - [ ] repos/{owner}/{repo}/topics/{topic}
   - [ ] repos/{owner}/{repo}/wiki/page/{pageName}
</details>

<details>
<summary><input checked disabled type="checkbox">Settings</input></summary>

 - **GET**
   - [X] /settings/api
   - [X] /settings/attachment
   - [X] /settings/repository
   - [X] /settings/ui

</details>

<details>
<summary><input disabled type="checkbox">User</input></summary>

 - **GET**
   - [X] /user
   - [ ] /user/applications/oauth2
   - [ ] /user/applications/oauth2/{id}
   - [X] /user/emails
   - [X] /user/followers
   - [X] /user/following
   - [ ] /user/following/{username}
   - [ ] /user/gpg_key_token
   - [X] /user/gpg_keys
   - [X] /user/gpg_keys/{id}
   - [X] /user/keys
   - [X] /user/keys/{id}
   - [X] /user/repos
   - [X] /user/settings
   - [ ] /user/starred
   - [ ] /user/starred/{owner}/{repo}
   - [ ] /user/stopwatches
   - [ ] /user/subscriptions
   - [ ] /user/teams
   - [ ] /user/times
   - [X] /users/search
   - [X] /users/{username}
   - [X] /users/{username}/followers
   - [X] /users/{username}/following
   - [ ] /users/{username}/following/{target}
   - [X] /users/{username}/gpg_keys
   - [ ] /users/{username}/heatmap
   - [X] /users/{username}/keys
   - [X] /users/{username}/repos
   - [ ] /users/{username}/starred
   - [ ] /users/{username}/subscriptions
   - [ ] /users/{username}/tokens


 - **POST**

   - [ ] /user/applications/oauth2
   - [ ] /user/emails
   - [ ] /user/gpg_key_verify
   - [ ] /user/gpg_keys
   - [ ] /user/keys
   - [X] /user/repos
   - [ ] /users/{username}/tokens

 - **PATCH**
   - [ ] /user/applications/oauth2/{id}
   - [X] /user/settings
 - **PUT**
   - [ ] /user/following/{username}
   - [ ] /user/starred/{owner}/{repo}
 - **DELETE**
   - [ ] /user/applications/oauth2/{id}
   - [ ] /user/emails
   - [ ] /user/following/{username}
   - [ ] /user/gpg_keys/{id}
   - [ ] /user/keys/{id}
   - [ ] /user/starred/{owner}/{repo}
   - [ ] /users/{username}/tokens/{token}
</details>
