package cmd

import (
	"github.com/hyperjiang/convert/client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var aliyunCmd = &cobra.Command{
	Use:   "aliyun",
	Short: "Aliyun machine translator",
	Run: func(cmd *cobra.Command, args []string) {
		aliOpts := client.AliyunOptions{
			RegionID:     viper.GetString("ALI_REGION_ID"),
			AccessKeyID:  viper.GetString("ALI_ACCESS_KEY_ID"),
			AccessSecret: viper.GetString("ALI_ACCESS_SECRET"),
		}
		c, err := client.NewAliyunClient(aliOpts)
		if err != nil {
			exit(err)
		}

		if opts.listLang {
			c.ListSupportedLanguages()
			return
		}

		if opts.inputFile != "" {
			if err := translate(c, opts); err != nil {
				exit(err)
			}

			if opts.outputFile != "/dev/stdout" {
				log.Info().Msgf("translate successfully and save into %s", opts.outputFile)
			}
		} else {
			cmd.Usage()
		}
	},
}

func init() {
	aliyunCmd.PersistentFlags().StringVarP(&opts.inputFile, "input", "i", "", "the input file to be translated, must provide")
	aliyunCmd.PersistentFlags().StringVarP(&opts.outputFile, "output", "o", "/dev/stdout", "the output path to save translated file")
	aliyunCmd.PersistentFlags().StringVarP(&opts.sourceLang, "source", "s", "en", "source language")
	aliyunCmd.PersistentFlags().StringVarP(&opts.targetLang, "targe", "t", "en", "target language")
	aliyunCmd.PersistentFlags().BoolVarP(&opts.listLang, "languages", "l", false, "list available languages")
}
