[Skip to main content](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#main) [Skip to in-page navigation](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#side-doc-outline)

## The future is yours

Microsoft Build · May 19 – 22, 2025

Join developers and AI innovators to refactor your skills at Microsoft Build.


[Register now](https://aka.ms/MSBuild_T1_Learn_Home)

Dismiss alert

This browser is no longer supported.

Upgrade to Microsoft Edge to take advantage of the latest features, security updates, and technical support.

[Download Microsoft Edge](https://go.microsoft.com/fwlink/p/?LinkID=2092881) [More info about Internet Explorer and Microsoft Edge](https://learn.microsoft.com/en-us/lifecycle/faq/internet-explorer-microsoft-edge)

Table of contentsExit focus mode

[Read in English](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest "Read in English")Add

- Add to Collections
- Add to plan

Table of contents [Read in English](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest "Read in English")Add to CollectionsAdd to plan

* * *

#### Share via

[Facebook](https://www.facebook.com/sharer/sharer.php?u=https%3A%2F%2Flearn.microsoft.com%2Fen-us%2Fcli%2Fazure%2Famlfs%3Fview%3Dazure-cli-latest%26WT.mc_id%3Dfacebook) [x.com](https://twitter.com/intent/tweet?original_referer=https%3A%2F%2Flearn.microsoft.com%2Fen-us%2Fcli%2Fazure%2Famlfs%3Fview%3Dazure-cli-latest%26WT.mc_id%3Dtwitter&text=Today%20I%20completed%20%22az%20amlfs%20%7C%20Microsoft%20Learn%22!%20I%27m%20so%20proud%20to%20be%20celebrating%20this%20achievement%20and%20hope%20this%20inspires%20you%20to%20start%20your%20own%20%40MicrosoftLearn%20journey!&tw_p=tweetbutton&url=https%3A%2F%2Flearn.microsoft.com%2Fen-us%2Fcli%2Fazure%2Famlfs%3Fview%3Dazure-cli-latest%26WT.mc_id%3Dtwitter) [LinkedIn](https://www.linkedin.com/feed/?shareActive=true&text=Today%20I%20completed%20%22az%20amlfs%20%7C%20Microsoft%20Learn%22!%20I%27m%20so%20proud%20to%20be%20celebrating%20this%20achievement%20and%20hope%20this%20inspires%20you%20to%20start%20your%20own%20%40MicrosoftLearn%20journey!%0A%0D%0Ahttps%3A%2F%2Flearn.microsoft.com%2Fen-us%2Fcli%2Fazure%2Famlfs%3Fview%3Dazure-cli-latest%26WT.mc_id%3Dlinkedin) [Email](mailto:?subject=%5BShared%20Article%5D%20az%20amlfs%20%7C%20Microsoft%20Learn&body=Today%20I%20completed%20%22az%20amlfs%20%7C%20Microsoft%20Learn%22!%20I'm%20so%20proud%20to%20be%20celebrating%20this%20achievement%20and%20hope%20this%20inspires%20you%20to%20start%20your%20own%20%40MicrosoftLearn%20journey!%0A%0D%0Ahttps%3A%2F%2Flearn.microsoft.com%2Fen-us%2Fcli%2Fazure%2Famlfs%3Fview%3Dazure-cli-latest%26WT.mc_id%3Demail)

* * *

Print

Table of contents

# az amlfs

- Reference

Feedback

Note

This reference is part of the **amlfs** extension for the Azure CLI (version 2.49.0 or higher). The extension will automatically install the first time you run an **az amlfs** command. [Learn more](https://learn.microsoft.com/en-us/cli/azure/azure-cli-extensions-overview) about extensions.

Manage lustre file system.

[Section titled: Commands](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#commands)

## Commands

Expand table

| Name | Description | Type | Status |
| --- | --- | --- | --- |
| [az amlfs archive](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-archive) | Archive data from the AML file system. | Extension | GA |
| [az amlfs cancel-archive](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-cancel-archive) | Cancel archiving data from the AML file system. | Extension | GA |
| [az amlfs check-amlfs-subnet](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-check-amlfs-subnet) | Check that subnets will be valid for AML file system create calls. | Extension | GA |
| [az amlfs create](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-create) | Create an AML file system. | Extension | GA |
| [az amlfs delete](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-delete) | Delete an AML file system for deletion. | Extension | GA |
| [az amlfs get-subnets-size](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-get-subnets-size) | Get the number of available IP addresses needed for the AML file system information provided. | Extension | GA |
| [az amlfs list](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-list) | List all AML file systems the user has access to under a resource group. | Extension | GA |
| [az amlfs show](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-show) | Get an AML file system. | Extension | GA |
| [az amlfs update](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-update) | Update an AML file system. | Extension | GA |
| [az amlfs wait](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-wait) | Place the CLI in a waiting state until a condition is met. | Extension | GA |

[Section titled: az amlfs archive](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-archive)

## az amlfs archive

Archive data from the AML file system.

Copy

```lang-azurecli
az amlfs archive [--amlfs-name]
                 [--filesystem-path]
                 [--ids]
                 [--resource-group]
                 [--subscription]
```

[Section titled: Examples](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-archive-examples)

### Examples

Amlfs archive

Copy

Open Cloud Shell

```lang-azurecli
az amlfs archive --amlfs-name name -g rg
```

[Section titled: Optional Parameters](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-archive-optional-parameters)

### Optional Parameters

--amlfs-name

Name for the AML file system. Allows alphanumerics, underscores, and hyphens. Start and end with alphanumeric.

--filesystem-path

Lustre file system path to archive relative to the file system root. Specify '/' to archive all modified data.

Default value: /

--ids

One or more resource IDs (space-delimited). It should be a complete resource ID containing all information of 'Resource Id' arguments. You should provide either --ids or other 'Resource Id' arguments.

--resource-group -g

Name of resource group. You can configure the default group using `az configure --defaults group=<name>`.

--subscription

Name or ID of subscription. You can configure the default subscription using `az account set -s NAME_OR_ID`.

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

[Section titled: az amlfs cancel-archive](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-cancel-archive)

## az amlfs cancel-archive

Cancel archiving data from the AML file system.

Copy

```lang-azurecli
az amlfs cancel-archive [--amlfs-name]
                        [--ids]
                        [--resource-group]
                        [--subscription]
```

[Section titled: Examples](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-cancel-archive-examples)

### Examples

Amlfs cancel-archive

Copy

Open Cloud Shell

```lang-azurecli
az amlfs cancel-archive --amlfs-name name -g rg
```

[Section titled: Optional Parameters](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-cancel-archive-optional-parameters)

### Optional Parameters

--amlfs-name

Name for the AML file system. Allows alphanumerics, underscores, and hyphens. Start and end with alphanumeric.

--ids

One or more resource IDs (space-delimited). It should be a complete resource ID containing all information of 'Resource Id' arguments. You should provide either --ids or other 'Resource Id' arguments.

--resource-group -g

Name of resource group. You can configure the default group using `az configure --defaults group=<name>`.

--subscription

Name or ID of subscription. You can configure the default subscription using `az account set -s NAME_OR_ID`.

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

[Section titled: az amlfs check-amlfs-subnet](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-check-amlfs-subnet)

## az amlfs check-amlfs-subnet

Check that subnets will be valid for AML file system create calls.

Copy

```lang-azurecli
az amlfs check-amlfs-subnet [--filesystem-subnet]
                            [--location]
                            [--sku]
                            [--storage-capacity]
```

[Section titled: Examples](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-check-amlfs-subnet-examples)

### Examples

Amlfs check subnet

Copy

Open Cloud Shell

```lang-azurecli
az amlfs check-amlfs-subnet --filesystem-subnet subnet_id --sku AMLFS-Durable-Premium-250 --location eastus --storage-capacity-tb 16
```

[Section titled: Optional Parameters](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-check-amlfs-subnet-optional-parameters)

### Optional Parameters

--filesystem-subnet

Subnet used for managing the AML file system and for client-facing operations. This subnet should have at least a /24 subnet mask within the VNET's address space.

--location

Region that the AML file system will be created in.

--sku

SKU name for this resource.

--storage-capacity

The size of the AML file system, in TiB.

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

[Section titled: az amlfs create](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-create)

## az amlfs create

Create an AML file system.

Copy

```lang-azurecli
az amlfs create --aml-filesystem-name
                --resource-group
                [--encryption-setting]
                [--filesystem-subnet]
                [--hsm-settings]
                [--location]
                [--maintenance-window]
                [--mi-user-assigned]
                [--no-wait {0, 1, f, false, n, no, t, true, y, yes}]
                [--sku]
                [--storage-capacity]
                [--tags]
                [--zones]
```

[Section titled: Examples](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-create-examples)

### Examples

Create amlfs

Copy

Open Cloud Shell

```lang-azurecli
az amlfs create -n amlfs_name -g rg --sku AMLFS-Durable-Premium-250 --storage-capacity 16 --zones [1] --maintenance-window "{dayOfWeek:friday,timeOfDayUtc:'22:00'}" --filesystem-subnet subnet_id
```

[Section titled: Required Parameters](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-create-required-parameters)

### Required Parameters

--aml-filesystem-name --name -n

Name for the AML file system. Allows alphanumerics, underscores, and hyphens. Start and end with alphanumeric.

--resource-group -g

Name of resource group. You can configure the default group using `az configure --defaults group=<name>`.

[Section titled: Optional Parameters](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-create-optional-parameters)

### Optional Parameters

--encryption-setting

Specifies the location of the encryption key in Key Vault. Support shorthand-syntax, json-file and yaml-file. Try "??" to show more.

--filesystem-subnet

Subnet used for managing the AML file system and for client-facing operations. This subnet should have at least a /24 subnet mask within the VNET's address space.

--hsm-settings

Specifies HSM settings of the AML file system. Support shorthand-syntax, json-file and yaml-file. Try "??" to show more.

--location -l

The geo-location where the resource lives When not specified, the location of the resource group will be used.

--maintenance-window

Start time of a 30-minute weekly maintenance window. Support shorthand-syntax, json-file and yaml-file. Try "??" to show more.

--mi-user-assigned

Space separated resource IDs to add user-assigned identities. Support shorthand-syntax, json-file and yaml-file. Try "??" to show more.

--no-wait

Do not wait for the long-running operation to finish.

Accepted values: 0, 1, f, false, n, no, t, true, y, yes

--sku

SKU name for this resource.

--storage-capacity

The size of the AML file system, in TiB. This might be rounded up.

--tags

Resource tags. Support shorthand-syntax, json-file and yaml-file. Try "??" to show more.

--zones

Availability zones for resources. This field should only contain a single element in the array. Support shorthand-syntax, json-file and yaml-file. Try "??" to show more.

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

[Section titled: az amlfs delete](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-delete)

## az amlfs delete

Delete an AML file system for deletion.

Copy

```lang-azurecli
az amlfs delete [--aml-filesystem-name]
                [--ids]
                [--no-wait {0, 1, f, false, n, no, t, true, y, yes}]
                [--resource-group]
                [--subscription]
                [--yes]
```

[Section titled: Examples](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-delete-examples)

### Examples

Delete amlfs

Copy

Open Cloud Shell

```lang-azurecli
az amlfs delete -n amlfs_name -g rg
```

[Section titled: Optional Parameters](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-delete-optional-parameters)

### Optional Parameters

--aml-filesystem-name --name -n

Name for the AML file system. Allows alphanumerics, underscores, and hyphens. Start and end with alphanumeric.

--ids

One or more resource IDs (space-delimited). It should be a complete resource ID containing all information of 'Resource Id' arguments. You should provide either --ids or other 'Resource Id' arguments.

--no-wait

Do not wait for the long-running operation to finish.

Accepted values: 0, 1, f, false, n, no, t, true, y, yes

--resource-group -g

Name of resource group. You can configure the default group using `az configure --defaults group=<name>`.

--subscription

Name or ID of subscription. You can configure the default subscription using `az account set -s NAME_OR_ID`.

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

[Section titled: az amlfs get-subnets-size](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-get-subnets-size)

## az amlfs get-subnets-size

Get the number of available IP addresses needed for the AML file system information provided.

Copy

```lang-azurecli
az amlfs get-subnets-size [--sku]
                          [--storage-capacity]
```

[Section titled: Examples](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-get-subnets-size-examples)

### Examples

Amlfs get subnet-size

Copy

Open Cloud Shell

```lang-azurecli
az amlfs get-subnets-size --sku AMLFS-Durable-Premium-250 --storage-capacity-tb 16
```

[Section titled: Optional Parameters](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-get-subnets-size-optional-parameters)

### Optional Parameters

--sku

SKU name for this resource.

--storage-capacity

The size of the AML file system, in TiB.

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

[Section titled: az amlfs list](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-list)

## az amlfs list

List all AML file systems the user has access to under a resource group.

Copy

```lang-azurecli
az amlfs list [--resource-group]
```

[Section titled: Examples](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-list-examples)

### Examples

List amlfs

Copy

Open Cloud Shell

```lang-azurecli
az amlfs list -g rg
```

[Section titled: Optional Parameters](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-list-optional-parameters)

### Optional Parameters

--resource-group -g

Name of resource group. You can configure the default group using `az configure --defaults group=<name>`.

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

[Section titled: az amlfs show](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-show)

## az amlfs show

Get an AML file system.

Copy

```lang-azurecli
az amlfs show [--aml-filesystem-name]
              [--ids]
              [--resource-group]
              [--subscription]
```

[Section titled: Examples](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-show-examples)

### Examples

Show amlfs

Copy

Open Cloud Shell

```lang-azurecli
az amlfs show -n name -g rg
```

[Section titled: Optional Parameters](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-show-optional-parameters)

### Optional Parameters

--aml-filesystem-name --name -n

Name for the AML file system. Allows alphanumerics, underscores, and hyphens. Start and end with alphanumeric.

--ids

One or more resource IDs (space-delimited). It should be a complete resource ID containing all information of 'Resource Id' arguments. You should provide either --ids or other 'Resource Id' arguments.

--resource-group -g

Name of resource group. You can configure the default group using `az configure --defaults group=<name>`.

--subscription

Name or ID of subscription. You can configure the default subscription using `az account set -s NAME_OR_ID`.

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

[Section titled: az amlfs update](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-update)

## az amlfs update

Update an AML file system.

Copy

```lang-azurecli
az amlfs update [--add]
                [--aml-filesystem-name]
                [--encryption-setting]
                [--force-string {0, 1, f, false, n, no, t, true, y, yes}]
                [--ids]
                [--maintenance-window]
                [--no-wait {0, 1, f, false, n, no, t, true, y, yes}]
                [--remove]
                [--resource-group]
                [--set]
                [--subscription]
                [--tags]
```

[Section titled: Examples](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-update-examples)

### Examples

Update amlfs

Copy

Open Cloud Shell

```lang-azurecli
az amlfs update -n name -g rg --tags "{tag:test}"
```

[Section titled: Optional Parameters](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-update-optional-parameters)

### Optional Parameters

--add

Add an object to a list of objects by specifying a path and key value pairs. Example: `--add property.listProperty <key=value, string or JSON string>`.

--aml-filesystem-name --name -n

Name for the AML file system. Allows alphanumerics, underscores, and hyphens. Start and end with alphanumeric.

--encryption-setting

Specifies the location of the encryption key in Key Vault. Support shorthand-syntax, json-file and yaml-file. Try "??" to show more.

--force-string

When using 'set' or 'add', preserve string literals instead of attempting to convert to JSON.

Accepted values: 0, 1, f, false, n, no, t, true, y, yes

--ids

One or more resource IDs (space-delimited). It should be a complete resource ID containing all information of 'Resource Id' arguments. You should provide either --ids or other 'Resource Id' arguments.

--maintenance-window

Start time of a 30-minute weekly maintenance window. Support shorthand-syntax, json-file and yaml-file. Try "??" to show more.

--no-wait

Do not wait for the long-running operation to finish.

Accepted values: 0, 1, f, false, n, no, t, true, y, yes

--remove

Remove a property or an element from a list. Example: `--remove property.list <indexToRemove>` OR `--remove propertyToRemove`.

--resource-group -g

Name of resource group. You can configure the default group using `az configure --defaults group=<name>`.

--set

Update an object by specifying a property path and value to set. Example: `--set property1.property2=<value>`.

--subscription

Name or ID of subscription. You can configure the default subscription using `az account set -s NAME_OR_ID`.

--tags

Resource tags. Support shorthand-syntax, json-file and yaml-file. Try "??" to show more.

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

[Section titled: az amlfs wait](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-wait)

## az amlfs wait

Place the CLI in a waiting state until a condition is met.

Copy

```lang-azurecli
az amlfs wait [--aml-filesystem-name]
              [--created]
              [--custom]
              [--deleted]
              [--exists]
              [--ids]
              [--interval]
              [--resource-group]
              [--subscription]
              [--timeout]
              [--updated]
```

[Section titled: Optional Parameters](https://learn.microsoft.com/en-us/cli/azure/amlfs?view=azure-cli-latest#az-amlfs-wait-optional-parameters)

### Optional Parameters

--aml-filesystem-name --name -n

Name for the AML file system. Allows alphanumerics, underscores, and hyphens. Start and end with alphanumeric.

--created

Wait until created with 'provisioningState' at 'Succeeded'.

Default value: False

--custom

Wait until the condition satisfies a custom JMESPath query. E.g. provisioningState!='InProgress', instanceView.statuses\[?code=='PowerState/running'\].

--deleted

Wait until deleted.

Default value: False

--exists

Wait until the resource exists.

Default value: False

--ids

One or more resource IDs (space-delimited). It should be a complete resource ID containing all information of 'Resource Id' arguments. You should provide either --ids or other 'Resource Id' arguments.

--interval

Polling interval in seconds.

Default value: 30

--resource-group -g

Name of resource group. You can configure the default group using `az configure --defaults group=<name>`.

--subscription

Name or ID of subscription. You can configure the default subscription using `az account set -s NAME_OR_ID`.

--timeout

Maximum wait in seconds.

Default value: 3600

--updated

Wait until updated with provisioningState at 'Succeeded'.

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

[Open a documentation issue](https://github.com/Azure/azure-cli/issues/new?template=docs_feedback.yml&pageUrl=https%3A%2F%2Flearn.microsoft.com%2Fen-us%2Fcli%2Fazure%2Famlfs%3Fview%3Dazure-cli-latest&pageQueryParams=%3Fview%3Dazure-cli-latest&contentSourceUrl=https%3A%2F%2Fgithub.com%2FMicrosoftDocs%2Fazure-docs-cli%2Fblob%2Fmain%2Fdocs-ref-autogen%2FLatest-version%2Flatest%2Famlfs.yml&documentVersionIndependentId=7263cf20-aa93-9d92-19c3-e0321203d400&platformId=3e7b02e2-51c7-09b8-cea4-fbbffc195e47&feedback=%0A%0A%5BEnter+feedback+here%5D%0A&author=%40mikefrobbins&metadata=*+ID%3A+4a297e65-1aad-56fb-1277-7409300d8701%0A*+PlatformId%3A+3e7b02e2-51c7-09b8-cea4-fbbffc195e47+%0A*+Service%3A+**azure-cli**&labels=needs-triage) [Provide product feedback](https://github.com/Azure/azure-cli/issues/new/choose)

## Additional resources
