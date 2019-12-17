// Code generated by openapi-cli-generator. DO NOT EDIT.
// See https://github.com/danielgtaylor/openapi-cli-generator

package main

import (
	"strings"

	"github.com/danielgtaylor/openapi-cli-generator/cli"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/h2non/gentleman.v2"
)

var openapiSubcommand bool

func openapiServers() []map[string]string {
	return []map[string]string{

		map[string]string{
			"description": "NOAA OneStop",
			"url":         "https://data.noaa.gov/onestop-search",
		},

		map[string]string{
			"description": "Development test server (uses test data)",
			"url":         "https://sciapps.colorado.edu/onestop-search",
		},
	}
}

// OpenapiGetCollection Get Collection Info
func OpenapiGetCollection(params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "getcollection"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/collection"

	req := cli.Client.Get().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// OpenapiHeadCollection Get Collection Info
func OpenapiHeadCollection(params *viper.Viper) (*gentleman.Response, interface{}, error) {
	handlerPath := "headcollection"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/collection"

	req := cli.Client.Head().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after
	}

	return resp, decoded, nil
}

// OpenapiGetCollectionById Collection by ID
func OpenapiGetCollectionById(paramId string, params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "getcollectionbyid"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/collection/{id}"
	url = strings.Replace(url, "{id}", paramId, 1)

	req := cli.Client.Get().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// OpenapiHeadCollectionById Collection by ID
func OpenapiHeadCollectionById(paramId string, params *viper.Viper) (*gentleman.Response, interface{}, error) {
	handlerPath := "headcollectionbyid"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/collection/{id}"
	url = strings.Replace(url, "{id}", paramId, 1)

	req := cli.Client.Head().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after
	}

	return resp, decoded, nil
}

// OpenapiGetFlattenedGranule Get Flattened Granule Info
func OpenapiGetFlattenedGranule(params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "getflattenedgranule"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/flattened-granule"

	req := cli.Client.Get().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// OpenapiHeadFlattenedGranule Get Flattened Granule Info
func OpenapiHeadFlattenedGranule(params *viper.Viper) (*gentleman.Response, interface{}, error) {
	handlerPath := "headflattenedgranule"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/flattened-granule"

	req := cli.Client.Head().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after
	}

	return resp, decoded, nil
}

// OpenapiGetFlattenedGranuleById Flattened Granule by ID
func OpenapiGetFlattenedGranuleById(paramId string, params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "getflattenedgranulebyid"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/flattened-granule/{id}"
	url = strings.Replace(url, "{id}", paramId, 1)

	req := cli.Client.Get().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// OpenapiHeadFlattenedGranuleById Flattened Granule by ID
func OpenapiHeadFlattenedGranuleById(paramId string, params *viper.Viper) (*gentleman.Response, interface{}, error) {
	handlerPath := "headflattenedgranulebyid"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/flattened-granule/{id}"
	url = strings.Replace(url, "{id}", paramId, 1)

	req := cli.Client.Head().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after
	}

	return resp, decoded, nil
}

// OpenapiHeadGranule Get Granule Info
func OpenapiHeadGranule(params *viper.Viper) (*gentleman.Response, interface{}, error) {
	handlerPath := "headgranule"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/granule"

	req := cli.Client.Head().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after
	}

	return resp, decoded, nil
}

// OpenapiGetGranule Get Granule Info
func OpenapiGetGranule(params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "getgranule"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/granule"

	req := cli.Client.Get().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// OpenapiGetGranuleById Granule by ID
func OpenapiGetGranuleById(paramId string, params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "getgranulebyid"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/granule/{id}"
	url = strings.Replace(url, "{id}", paramId, 1)

	req := cli.Client.Get().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// OpenapiHeadGranuleById Granule by ID
func OpenapiHeadGranuleById(paramId string, params *viper.Viper) (*gentleman.Response, interface{}, error) {
	handlerPath := "headgranulebyid"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/granule/{id}"
	url = strings.Replace(url, "{id}", paramId, 1)

	req := cli.Client.Head().URL(url)

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after
	}

	return resp, decoded, nil
}

// OpenapiSearchCollection Retrieve collection metadata
func OpenapiSearchCollection(params *viper.Viper, body string) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "searchcollection"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/search/collection"

	req := cli.Client.Post().URL(url)

	if body != "" {
		req = req.AddHeader("Content-Type", "application/json").BodyString(body)
	}

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// OpenapiSearchFlattenedGranule Retrieve flattened granule metadata
func OpenapiSearchFlattenedGranule(params *viper.Viper, body string) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "searchflattenedgranule"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/search/flattened-granule"

	req := cli.Client.Post().URL(url)

	if body != "" {
		req = req.AddHeader("Content-Type", "application/json").BodyString(body)
	}

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

// OpenapiSearchGranule Retrieve granule metadata
func OpenapiSearchGranule(params *viper.Viper, body string) (*gentleman.Response, map[string]interface{}, error) {
	handlerPath := "searchgranule"
	if openapiSubcommand {
		handlerPath = "openapi " + handlerPath
	}

	server := viper.GetString("server")
	if server == "" {
		server = openapiServers()[viper.GetInt("server-index")]["url"]
	}

	url := server + "/search/granule"

	req := cli.Client.Post().URL(url)

	if body != "" {
		req = req.AddHeader("Content-Type", "application/json").BodyString(body)
	}

	cli.HandleBefore(handlerPath, params, req)

	resp, err := req.Do()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Request failed")
	}

	var decoded map[string]interface{}

	if resp.StatusCode < 400 {
		if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
			return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
		}
	} else {
		return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
	}

	after := cli.HandleAfter(handlerPath, params, resp, decoded)
	if after != nil {
		decoded = after.(map[string]interface{})
	}

	return resp, decoded, nil
}

func openapiRegister(subcommand bool) {
	root := cli.Root

	if subcommand {
		root = &cobra.Command{
			Use:   "openapi",
			Short: "OneStop Search API",
			Long:  cli.Markdown("Search Collections and Granules! More information on search request and responses available at [Search API Requests](https://github.com/cedardevs/onestop/wiki/OneStop-Search-API-Requests) and [Search API Responses](https://github.com/cedardevs/onestop/wiki/OneStop-Search-API-Responses)."),
		}
		openapiSubcommand = true
	} else {
		cli.Root.Short = "OneStop Search API"
		cli.Root.Long = cli.Markdown("Search Collections and Granules! More information on search request and responses available at [Search API Requests](https://github.com/cedardevs/onestop/wiki/OneStop-Search-API-Requests) and [Search API Responses](https://github.com/cedardevs/onestop/wiki/OneStop-Search-API-Responses).")
	}

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "getcollection",
			Short:   "Get Collection Info",
			Long:    cli.Markdown("Get the total number of collections available."),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := OpenapiGetCollection(params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "headcollection",
			Short:   "Get Collection Info",
			Long:    cli.Markdown(""),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := OpenapiHeadCollection(params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "getcollectionbyid id",
			Short:   "Collection by ID",
			Long:    cli.Markdown(""),
			Example: examples,
			Args:    cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := OpenapiGetCollectionById(args[0], params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "headcollectionbyid id",
			Short:   "Collection by ID",
			Long:    cli.Markdown(""),
			Example: examples,
			Args:    cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := OpenapiHeadCollectionById(args[0], params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "getflattenedgranule",
			Short:   "Get Flattened Granule Info",
			Long:    cli.Markdown("Get the total number of flattened granules available."),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := OpenapiGetFlattenedGranule(params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "headflattenedgranule",
			Short:   "Get Flattened Granule Info",
			Long:    cli.Markdown(""),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := OpenapiHeadFlattenedGranule(params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "getflattenedgranulebyid id",
			Short:   "Flattened Granule by ID",
			Long:    cli.Markdown(""),
			Example: examples,
			Args:    cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := OpenapiGetFlattenedGranuleById(args[0], params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "headflattenedgranulebyid id",
			Short:   "Flattened Granule by ID",
			Long:    cli.Markdown(""),
			Example: examples,
			Args:    cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := OpenapiHeadFlattenedGranuleById(args[0], params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "headgranule",
			Short:   "Get Granule Info",
			Long:    cli.Markdown(""),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := OpenapiHeadGranule(params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "getgranule",
			Short:   "Get Granule Info",
			Long:    cli.Markdown("Get the total number of granules available."),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := OpenapiGetGranule(params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "getgranulebyid id",
			Short:   "Granule by ID",
			Long:    cli.Markdown(""),
			Example: examples,
			Args:    cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := OpenapiGetGranuleById(args[0], params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "headgranulebyid id",
			Short:   "Granule by ID",
			Long:    cli.Markdown(""),
			Example: examples,
			Args:    cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

				_, decoded, err := OpenapiHeadGranuleById(args[0], params)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		examples += "  " + cli.Root.CommandPath() + " searchcollection facets: true, page{max: 20, offset: 0}, queries[]{type: queryText, value: climate}\n"

		cmd := &cobra.Command{
			Use:     "searchcollection",
			Short:   "Retrieve collection metadata",
			Long:    cli.Markdown("Retrieve collection metadata records matching the text query string, spatial, and/or temporal filter.\n## Request Schema (application/json)\n\nadditionalProperties: false\ndescription: The shape of a search request body that can be sent to the OneStop API\n  to execute a search against available metadata. Collections are returned by default\n  unless a collection filter is included, resulting in granules being returned.\nproperties:\n  facets:\n    default: false\n    description: Flag to request counts of results by GCMD keywords in addition to\n      results.\n    type: boolean\n  filters:\n    items:\n      anyOf:\n      - $ref: '#/components/schemas/dateTimeFilter'\n      - $ref: '#/components/schemas/yearFilter'\n      - $ref: '#/components/schemas/facetFilter'\n      - $ref: '#/components/schemas/geometryFilter'\n      - $ref: '#/components/schemas/collectionFilter'\n      - $ref: '#/components/schemas/excludeGlobalFilter'\n    type: array\n  page:\n    $ref: '#/components/schemas/page'\n  queries:\n    items:\n      oneOf:\n      - $ref: '#/components/schemas/textQuery'\n    type: array\n  summary:\n    default: true\n    description: Flag to request summary of search results instead of full set of\n      attributes.\n    type: boolean\ntitle: Search Request\ntype: object\n"),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {
				body, err := cli.GetBody("application/json", args[0:])
				if err != nil {
					log.Fatal().Err(err).Msg("Unable to get body")
				}

				_, decoded, err := OpenapiSearchCollection(params, body)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "searchflattenedgranule",
			Short:   "Retrieve flattened granule metadata",
			Long:    cli.Markdown("Retrieve flattened granule metadata records matching the text query string, spatial, and/or temporal filter.\n## Request Schema (application/json)\n\nadditionalProperties: false\ndescription: The shape of a search request body that can be sent to the OneStop API\n  to execute a search against available metadata. Collections are returned by default\n  unless a collection filter is included, resulting in granules being returned.\nproperties:\n  facets:\n    default: false\n    description: Flag to request counts of results by GCMD keywords in addition to\n      results.\n    type: boolean\n  filters:\n    items:\n      anyOf:\n      - $ref: '#/components/schemas/dateTimeFilter'\n      - $ref: '#/components/schemas/yearFilter'\n      - $ref: '#/components/schemas/facetFilter'\n      - $ref: '#/components/schemas/geometryFilter'\n      - $ref: '#/components/schemas/collectionFilter'\n      - $ref: '#/components/schemas/excludeGlobalFilter'\n    type: array\n  page:\n    $ref: '#/components/schemas/page'\n  queries:\n    items:\n      oneOf:\n      - $ref: '#/components/schemas/textQuery'\n    type: array\n  summary:\n    default: true\n    description: Flag to request summary of search results instead of full set of\n      attributes.\n    type: boolean\ntitle: Search Request\ntype: object\n"),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {
				body, err := cli.GetBody("application/json", args[0:])
				if err != nil {
					log.Fatal().Err(err).Msg("Unable to get body")
				}

				_, decoded, err := OpenapiSearchFlattenedGranule(params, body)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

	func() {
		params := viper.New()

		var examples string

		cmd := &cobra.Command{
			Use:     "searchgranule",
			Short:   "Retrieve granule metadata",
			Long:    cli.Markdown("Retrieve granule metadata records matching the text query string, spatial, and/or temporal filter.\n## Request Schema (application/json)\n\nadditionalProperties: false\ndescription: The shape of a search request body that can be sent to the OneStop API\n  to execute a search against available metadata. Collections are returned by default\n  unless a collection filter is included, resulting in granules being returned.\nproperties:\n  facets:\n    default: false\n    description: Flag to request counts of results by GCMD keywords in addition to\n      results.\n    type: boolean\n  filters:\n    items:\n      anyOf:\n      - $ref: '#/components/schemas/dateTimeFilter'\n      - $ref: '#/components/schemas/yearFilter'\n      - $ref: '#/components/schemas/facetFilter'\n      - $ref: '#/components/schemas/geometryFilter'\n      - $ref: '#/components/schemas/collectionFilter'\n      - $ref: '#/components/schemas/excludeGlobalFilter'\n    type: array\n  page:\n    $ref: '#/components/schemas/page'\n  queries:\n    items:\n      oneOf:\n      - $ref: '#/components/schemas/textQuery'\n    type: array\n  summary:\n    default: true\n    description: Flag to request summary of search results instead of full set of\n      attributes.\n    type: boolean\ntitle: Search Request\ntype: object\n"),
			Example: examples,
			Args:    cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {
				body, err := cli.GetBody("application/json", args[0:])
				if err != nil {
					log.Fatal().Err(err).Msg("Unable to get body")
				}

				_, decoded, err := OpenapiSearchGranule(params, body)
				if err != nil {
					log.Fatal().Err(err).Msg("Error calling operation")
				}

				if err := cli.Formatter.Format(decoded); err != nil {
					log.Fatal().Err(err).Msg("Formatting failed")
				}

			},
		}
		root.AddCommand(cmd)

		cli.SetCustomFlags(cmd)

		if cmd.Flags().HasFlags() {
			params.BindPFlags(cmd.Flags())
		}

	}()

}
