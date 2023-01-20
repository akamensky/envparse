package envparse

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type envMap map[string]string

func newEnvMap(prefix string, envList []string) envMap {
	env := make(map[string]string)
	{
		for _, envVar := range envList {
			envVarParts := strings.SplitN(envVar, "=", 2)
			if len(envVarParts) != 2 {
				// Skip invalid envVar
				continue
			}
			if strings.HasPrefix(envVarParts[0], prefix) {
				env[envVarParts[0]] = envVarParts[1]
			}
		}
	}

	return env
}

func (env envMap) Exists(key string) bool {
	if _, ok := env[key]; ok {
		return true
	}

	return false
}

func (env envMap) PrefixExists(prefix string) bool {
	for k := range env {
		if strings.HasPrefix(k, prefix) {
			return true
		}
	}

	return false
}

func (env envMap) Get(key string) string {
	return env[key]
}

func (env envMap) GetPrefix(prefix string) envMap {
	result := make(map[string]string)

	prefix = fmt.Sprintf("%s_", strings.TrimRight(prefix, "_"))

	for k, v := range env {
		if strings.HasPrefix(k, prefix) {
			result[strings.TrimPrefix(k, prefix)] = v
		}
	}

	return result
}

func (env envMap) GetSlicePrefixes() []string {
	keys := make(map[int]string)

	re := regexp.MustCompile(`^[0-9]+$`)
	for k := range env {
		kPart := strings.Split(k, "_")[0]
		if re.MatchString(kPart) {
			numValue, _ := strconv.Atoi(kPart)
			keys[numValue] = kPart
		}
	}

	result := make([]string, 0)
	order := make([]int, 0)

	for i := range keys {
		order = append(order, i)
	}

	sort.Ints(order)

	for _, i := range order {
		result = append(result, keys[i])
	}

	return result
}
