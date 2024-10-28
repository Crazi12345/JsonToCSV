package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type User struct {
    Hireable      *bool      `json:"hirable"`
    PublicRepos   *int       `json:"public_repos"`
    IsSuspicious *bool      `json:"is_suspicious"`
    UpdatedAt    *string `json:"updated_at"`
    ID           *int64     `json:"id"`
    Blog         *string    `json:"blog"`
    Followers    *int       `json:"followers"`
    Location     *string    `json:"location"`
    FollowerList *[]interface{} `json:"follower_list"`
    Type         *string    `json:"type"`
    CommitList   *[]Commit   `json:"commit_list"`
    Bio         *string    `json:"bio"`
    Commits      *int       `json:"commits"` // This field was in your JSON
    Company     *string    `json:"company"`
    FollowingList *[]interface{} `json:"following_list"`
    PublicGists  *int       `json:"public_gists"`
    Name         *string    `json:"name"`
    CreatedAt    *string `json:"created_at"`
    Email        *string    `json:"email"`
    Following    *int       `json:"following"`
    Login        *string    `json:"login"`
    RepoList     *[]Repo     `json:"repo_list"`
}

type Commit struct {
    RepoID          *int64     `json:"repo_id"`
    RepoOwnerID     *int64     `json:"repo_owner_id"`
    CommitAt        *string `json:"commit_at"`
    CommitterID     *int64     `json:"committer_id"`
    Message         *string    `json:"message"`
    RepoDescription *string    `json:"repo_description"`
    GenerateAt      *string `json:"generate_at"`
    AuthorID        *int64     `json:"author_id"`
    RepoName        *string    `json:"repo_name"`
}

type Repo struct {
    Fork           *bool      `json:"fork"`
    License        *string    `json:"license"`
    HasWiki        *bool      `json:"has_wiki"`
    Description    *string    `json:"description"`
    Language       *string    `json:"language"`
    DefaultBranch  *string    `json:"default_branch"`
    CreatedAt      *string `json:"created_at"`
    ForksCount     *int       `json:"forks_count"`
    UpdatedAt      *string `json:"updated_at"`
    PushedAt       *string `json:"pushed_at"`
    FullName       *string    `json:"full_name"`
    OpenIssues     *int       `json:"open_issues"`
    StargazersCount *int       `json:"stargazers_count"`
    OwnerID        *int64     `json:"owner_id"`
    ID             *int64     `json:"id"`
    Size           *int       `json:"size"`
}
func check(e error, message string) {
	if e != nil {
		log.Fatal(message)
	}
}

func popLine(f *os.File) ([]byte, error) {
	fi, err := f.Stat()
	check(err, "file statistics failed")
	buf := bytes.NewBuffer(make([]byte, 0, fi.Size()))

	_, err = f.Seek(0, io.SeekStart)
	check(err, "seek failed")

	_, err = io.Copy(buf, f)
	check(err, "copy failed")

	line, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}

	_, err = f.Seek(0, io.SeekStart)
	check(err, "seek failed")

	nw, err := io.Copy(f, buf)
	check(err, "copy failed")

	err = f.Truncate(nw)
	check(err, "truncating failed")

	err = f.Sync()
	check(err, "syncing failed")

	_, err = f.Seek(0, io.SeekStart)
	check(err, "seek failed")

	return line, nil
}
func UserToCSV(user User) {

}
func main() {
	file, err := os.OpenFile("/home/tired_atlas/Desktop/test.json", os.O_RDWR, 0644)
	check(err, "Could not open file")
	defer file.Close()
    for i := 0; i<100; i++ {
	data, err := popLine(file)
	check(err, "could not read file")

	var user User

	error := json.Unmarshal([]byte(data), &user)
	check(error, "failed to decode json")
        fmt.Println(i)
    }
}
