[Skip to main content](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#main) [Skip to in-page navigation](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#side-doc-outline)

## The future is yours

Microsoft Build · May 19 – 22, 2025

Join developers and AI innovators to refactor your skills at Microsoft Build.


[Register now](https://aka.ms/MSBuild_T1_Learn_Home)

Dismiss alert

This browser is no longer supported.

Upgrade to Microsoft Edge to take advantage of the latest features, security updates, and technical support.

[Download Microsoft Edge](https://go.microsoft.com/fwlink/p/?LinkID=2092881) [More info about Internet Explorer and Microsoft Edge](https://learn.microsoft.com/en-us/lifecycle/faq/internet-explorer-microsoft-edge)

Table of contentsExit focus mode

[Read in English](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest "Read in English")Add

- Add to Collections
- Add to plan

Table of contents [Read in English](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest "Read in English")Add to CollectionsAdd to plan

* * *

#### Share via

[Facebook](https://www.facebook.com/sharer/sharer.php?u=https%3A%2F%2Flearn.microsoft.com%2Fen-us%2Fcli%2Fazure%2Falias%3Fview%3Dazure-cli-latest%26WT.mc_id%3Dfacebook) [x.com](https://twitter.com/intent/tweet?original_referer=https%3A%2F%2Flearn.microsoft.com%2Fen-us%2Fcli%2Fazure%2Falias%3Fview%3Dazure-cli-latest%26WT.mc_id%3Dtwitter&text=Today%20I%20completed%20%22az%20alias%20%7C%20Microsoft%20Learn%22!%20I%27m%20so%20proud%20to%20be%20celebrating%20this%20achievement%20and%20hope%20this%20inspires%20you%20to%20start%20your%20own%20%40MicrosoftLearn%20journey!&tw_p=tweetbutton&url=https%3A%2F%2Flearn.microsoft.com%2Fen-us%2Fcli%2Fazure%2Falias%3Fview%3Dazure-cli-latest%26WT.mc_id%3Dtwitter) [LinkedIn](https://www.linkedin.com/feed/?shareActive=true&text=Today%20I%20completed%20%22az%20alias%20%7C%20Microsoft%20Learn%22!%20I%27m%20so%20proud%20to%20be%20celebrating%20this%20achievement%20and%20hope%20this%20inspires%20you%20to%20start%20your%20own%20%40MicrosoftLearn%20journey!%0A%0D%0Ahttps%3A%2F%2Flearn.microsoft.com%2Fen-us%2Fcli%2Fazure%2Falias%3Fview%3Dazure-cli-latest%26WT.mc_id%3Dlinkedin) [Email](mailto:?subject=%5BShared%20Article%5D%20az%20alias%20%7C%20Microsoft%20Learn&body=Today%20I%20completed%20%22az%20alias%20%7C%20Microsoft%20Learn%22!%20I'm%20so%20proud%20to%20be%20celebrating%20this%20achievement%20and%20hope%20this%20inspires%20you%20to%20start%20your%20own%20%40MicrosoftLearn%20journey!%0A%0D%0Ahttps%3A%2F%2Flearn.microsoft.com%2Fen-us%2Fcli%2Fazure%2Falias%3Fview%3Dazure-cli-latest%26WT.mc_id%3Demail)

* * *

Print

Table of contents

# az alias

- Reference

Feedback

Note

This reference is part of the **alias** extension for the Azure CLI (version 2.0.31.dev0 or higher). The extension will automatically install the first time you run an **az alias** command. [Learn more](https://learn.microsoft.com/en-us/cli/azure/azure-cli-extensions-overview) about extensions.

Manage Azure CLI Aliases.

[Section titled: Commands](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#commands)

## Commands

Expand table

| Name | Description | Type | Status |
| --- | --- | --- | --- |
| [az alias create](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-create) | Create an alias. | Extension | GA |
| [az alias export](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-export) | Export all registered aliases to a given path, as an INI configuration file. If no export path is specified, the alias configuration file is exported to the current working directory. | Extension | GA |
| [az alias import](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-import) | Import aliases from an INI configuration file or an URL. | Extension | GA |
| [az alias list](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-list) | List the registered aliases. | Extension | GA |
| [az alias remove](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-remove) | Remove one or more aliases. Aliases to be removed are space-delimited. | Extension | GA |
| [az alias remove-all](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-remove-all) | Remove all registered aliases. | Extension | GA |

[Section titled: az alias create](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-create)

## az alias create

Create an alias.

Copy

```lang-azurecli
az alias create --command
                --name
```

[Section titled: Examples](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-create-examples)

### Examples

Create simple alias commands.

Copy

Open Cloud Shell

```lang-azurecli
az alias create --name rg --command group

az alias create --name ls --command list
```

Create a complex alias.

Copy

Open Cloud Shell

```lang-azurecli
az alias create --name list-vm --command 'vm list --resource-group myResourceGroup'
```

Create an alias command with arguments.

Copy

Open Cloud Shell

```lang-azurecli
az alias create --name 'list-vm {{ resource_group }}' \
  --command 'vm list --resource-group {{ resource_group }}'
```

Process arguments using Jinja2 templates.

Copy

Open Cloud Shell

```lang-azurecli
az alias create --name 'storage-ls {{ url }}' \
  --command 'storage blob list
    --account-name {{ url.replace("https://", "").split(".")[0] }}
    --container-name {{ url.replace("https://", "").split("/")[1] }}'
```

[Section titled: Required Parameters](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-create-required-parameters)

### Required Parameters

--command -c

The command that the alias points to.

--name -n

The name of the alias.

Global Parameters

--debug

Increase logging verbosity to show all debug logs.

--help -h

Show this help message and exit.

--only-show-errors

Only show errors, suppressing warnings.

--output -o

Output format.

Accepted values: json, jsonc, none, table, tsv, yaml, yamlc

Default value: json

--query

JMESPath query string. See [http://jmespath.org/](http://jmespath.org/) for more information and examples.

--subscription

Name or ID of subscription. You can configure the default subscription using `az account set -s NAME_OR_ID`.

--verbose

Increase logging verbosity. Use --debug for full debug logs.

[Section titled: az alias export](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-export)

## az alias export

Export all registered aliases to a given path, as an INI configuration file. If no export path is specified, the alias configuration file is exported to the current working directory.

Copy

```lang-azurecli
az alias export [--exclude]
                [--path]
```

[Section titled: Optional Parameters](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-export-optional-parameters)

### Optional Parameters

--exclude -e

Space-separated aliases excluded from export.

--path -p

The path of the alias configuration file to export to.

Global Parameters

--debug

Increase logging verbosity to show all debug logs.

--help -h

Show this help message and exit.

--only-show-errors

Only show errors, suppressing warnings.

--output -o

Output format.

Accepted values: json, jsonc, none, table, tsv, yaml, yamlc

Default value: json

--query

JMESPath query string. See [http://jmespath.org/](http://jmespath.org/) for more information and examples.

--subscription

Name or ID of subscription. You can configure the default subscription using `az account set -s NAME_OR_ID`.

--verbose

Increase logging verbosity. Use --debug for full debug logs.

[Section titled: az alias import](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-import)

## az alias import

Import aliases from an INI configuration file or an URL.

Copy

```lang-azurecli
az alias import --source
```

[Section titled: Required Parameters](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-import-required-parameters)

### Required Parameters

--source -s

The source of the aliases to import from.

Global Parameters

--debug

Increase logging verbosity to show all debug logs.

--help -h

Show this help message and exit.

--only-show-errors

Only show errors, suppressing warnings.

--output -o

Output format.

Accepted values: json, jsonc, none, table, tsv, yaml, yamlc

Default value: json

--query

JMESPath query string. See [http://jmespath.org/](http://jmespath.org/) for more information and examples.

--subscription

Name or ID of subscription. You can configure the default subscription using `az account set -s NAME_OR_ID`.

--verbose

Increase logging verbosity. Use --debug for full debug logs.

[Section titled: az alias list](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-list)

## az alias list

List the registered aliases.

Copy

```lang-azurecli
az alias list
```

Global Parameters

--debug

Increase logging verbosity to show all debug logs.

--help -h

Show this help message and exit.

--only-show-errors

Only show errors, suppressing warnings.

--output -o

Output format.

Accepted values: json, jsonc, none, table, tsv, yaml, yamlc

Default value: json

--query

JMESPath query string. See [http://jmespath.org/](http://jmespath.org/) for more information and examples.

--subscription

Name or ID of subscription. You can configure the default subscription using `az account set -s NAME_OR_ID`.

--verbose

Increase logging verbosity. Use --debug for full debug logs.

[Section titled: az alias remove](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-remove)

## az alias remove

Remove one or more aliases. Aliases to be removed are space-delimited.

Copy

```lang-azurecli
az alias remove --name
```

[Section titled: Required Parameters](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-remove-required-parameters)

### Required Parameters

--name -n

Space-separated aliases.

Global Parameters

--debug

Increase logging verbosity to show all debug logs.

--help -h

Show this help message and exit.

--only-show-errors

Only show errors, suppressing warnings.

--output -o

Output format.

Accepted values: json, jsonc, none, table, tsv, yaml, yamlc

Default value: json

--query

JMESPath query string. See [http://jmespath.org/](http://jmespath.org/) for more information and examples.

--subscription

Name or ID of subscription. You can configure the default subscription using `az account set -s NAME_OR_ID`.

--verbose

Increase logging verbosity. Use --debug for full debug logs.

[Section titled: az alias remove-all](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-remove-all)

## az alias remove-all

Remove all registered aliases.

Copy

```lang-azurecli
az alias remove-all [--yes]
```

[Section titled: Optional Parameters](https://learn.microsoft.com/en-us/cli/azure/alias?view=azure-cli-latest#az-alias-remove-all-optional-parameters)

### Optional Parameters

--yes -y

Do not prompt for confirmation.

Default value: False

Global Parameters

--debug

Increase logging verbosity to show all debug logs.

--help -h

Show this help message and exit.

--only-show-errors

Only show errors, suppressing warnings.

--output -o

Output format.

Accepted values: json, jsonc, none, table, tsv, yaml, yamlc

Default value: json

--query

JMESPath query string. See [http://jmespath.org/](http://jmespath.org/) for more information and examples.

--subscription

Name or ID of subscription. You can configure the default subscription using `az account set -s NAME_OR_ID`.

--verbose

Increase logging verbosity. Use --debug for full debug logs.

Collaborate with us on GitHub

The source for this content can be found on GitHub, where you can also create and review issues and pull requests. For more information, see [our contributor guide](https://learn.microsoft.com/contribute/content/how-to-write-overview).


![](https://learn.microsoft.com/media/logos/logo_azure.svg)![](https://learn.microsoft.com/media/logos/logo_azure.svg)

Azure CLI feedback

Azure CLI is an open source project. Select a link to provide feedback:

[Open a documentation issue](https://github.com/Azure/azure-cli/issues/new?template=docs_feedback.yml&pageUrl=https%3A%2F%2Flearn.microsoft.com%2Fen-us%2Fcli%2Fazure%2Falias%3Fview%3Dazure-cli-latest&pageQueryParams=%3Fview%3Dazure-cli-latest&contentSourceUrl=https%3A%2F%2Fgithub.com%2FMicrosoftDocs%2Fazure-docs-cli%2Fblob%2Fmain%2Fdocs-ref-autogen%2FLatest-version%2Flatest%2Falias.yml&documentVersionIndependentId=070d580e-9282-6b95-2b8a-1a0f406e8e0c&platformId=53789d7f-e5d7-6edb-c2a9-6aa0d5c54a5e&feedback=%0A%0A%5BEnter+feedback+here%5D%0A&author=%40mikefrobbins&metadata=*+ID%3A+f333637a-c06e-2a9c-5df8-6f22b8502be5%0A*+PlatformId%3A+53789d7f-e5d7-6edb-c2a9-6aa0d5c54a5e+%0A*+Service%3A+**azure-cli**&labels=needs-triage) [Provide product feedback](https://github.com/Azure/azure-cli/issues/new/choose)

## Additional resources
