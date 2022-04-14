package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	git_user_names := []string{}
	git_user_names = append(git_user_names, user.Username)
	git_user_names = append(git_user_names, get_git_user_name())

	ghq_root_paths, err := exec.Command("ghq", "root", "-all").Output()
	if err != nil {
		log.Fatal(err)
	}
	grps := strings.Split(string(ghq_root_paths), "\n")
	for _, grp := range grps {
		fmt.Println(grp)
		if strings.HasSuffix(grp, "ghq") {
			for _, git_user_name := range git_user_names {
				github_repos_with_user := grp + "/github.com/" + git_user_name
				repositories, err := ioutil.ReadDir(github_repos_with_user)
				if err != nil {
					log.Fatal(err)
				}
				for _, repo := range repositories {
					if repo.IsDir() {
						git_repo_path := github_repos_with_user + "/" + repo.Name()
						log, err := get_recent_git_logs(git_repo_path)
						if err == nil && log != "" {
							fmt.Println(git_repo_path)
							fmt.Println(log)
						}
					}
				}
			}
		}
	}
}

func get_recent_git_logs(path string) (string, error) {
	out, err := exec.Command("git", "-C", path, "log", "--oneline", "--since=\"1 days ago\"").Output()
	if err != nil {
		return "", err

	}
	return string(out), nil
}

func get_git_user_name() string {
	name, err := exec.Command("git", "config", "--get", "user.name").Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimRight(string(name), "\n")
}
