package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func getEnv(k string, def string) (string) {
	value := def
	if v := os.Getenv(k); v != "" {
		value = v
	}
	return value
}

func get_url_text(url string) (string) {
	resp, err := http.Get(url)
	check(err)
		
	defer resp.Body.Close()
	
	data, err := ioutil.ReadAll(resp.Body)
	check(err)
	
	return string(data)
}

// Gets a key from AWS metadata
func get_meta(param string) (string) {
	u, err := url.Parse(getEnv("METADATA_URL", "http://169.254.169.254"))
	check(err)
	u.Path = path.Join(u.Path, "/latest/meta-data/", param)
	return get_url_text(u.String())
}

// foo-bar --> FOO_BAR
func as_env_var(key string, val string) (string) {
	KEY := path.Base(strings.Replace(strings.ToUpper(key), "-", "_", -1))
	return fmt.Sprintf("%s=\"%s\"", KEY, val)
}

func write_vars(f string, values []string) (error) {
	file, err := os.Create(f)
	check(err)
	
	defer file.Close()
	
	writer := bufio.NewWriter(file)
	for _, line := range values {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}

func main() {
	all_params := []string{
		"ami-id", 
		"ami-launch-index",
		"ami-manifest-path",
		"hostname",
		"instance-action",
		"instance-id",
		"instance-type",
		"local-hostname",
		"local-ipv4",
		"mac",
		"placement/availability-zone",
		"profile",
		"public-hostname",
		"public-ipv4",
		"reservation-id",
		"security-groups",
	}

	values := []string{}
	
	for _, p := range all_params {
		v := as_env_var(p, get_meta(p))
		values = append(values, v)
		fmt.Println(fmt.Sprintf("export %s", v))
	}
	
	if f := getEnv("AWS_ENVIRONMENT_FILE", ""); f != "" {
		err := write_vars(f, values)
		check(err)
	}
}
