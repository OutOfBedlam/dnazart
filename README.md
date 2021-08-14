# Download Azure Pipeline Artifacts

`dnazart` requires PAT from environment variable

```bash
$ export AZURE_DEVOPS_EXT_PAT='xxxxxxxx'
```

## Usage

```bash
$ dnazart --help

Usage: dnazart COMMAND [arg...]

download artifact files from Azure DevOps Pipelines
                  
Commands:         
  version         show version
  download        download artifacts
  hist, history   list builds history
                  
Run 'dnazart COMMAND --help' for more information on a command.
```

## List build history

```bash
$ dnazart list <organization> <project>
```

- output example

```
BuildId  BuildNumber                     Status     FinishTime
121      your_build_master_20210813.2   ✓ Success  2021-08-13T08:19:28.9561614Z
120      your_build_master_20210813.1   ✓ Success  2021-08-13T06:22:01.6570432Z
119      your_build_master_20210812.2   ✓ Success  2021-08-12T08:15:38.882945Z
118      your_build_master_20210812.1   ✓ Success  2021-08-12T01:16:16.8851851Z
117      your_build_master_20210811.6   ✓ Success  2021-08-11T05:00:57.8963808Z
116      your_build_master_20210811.5   ✓ Success  2021-08-11T03:25:23.8824055Z
115      your_build_master_20210811.4   𐄂 Fail     2021-08-11T01:21:50.6766667Z
114      your_build_master_20210811.3   𐄂 Fail     2021-08-11T00:22:18.8291786Z
```

- options
```
Usage: dnazart hist [-n=<limit>] ORG PRJ

list builds history
               
Arguments:     
  ORG          organization
  PRJ          project
               
Options:       
  -n, --num    limit number of output, 0: no limit (default 0)
```

## Download artifact from build

- download artifact from latest build.

```bash
$ dnazart download <organization> <project>
```


- download from a specific build

```bash
$ dnazart download <organization> <project> <BuildId>
```

- options

```
Usage: dnazart download [-d] [-f] [-o=<output path>] ORG PRJ [BUILDID]

download artifacts
                
Arguments:      
  ORG           organization
  PRJ           project
  BUILDID       build id (default "latest")
                
Options:        
  -d, --dry     dry-run, print download url without actual downloading
  -f, --force   overwrite if output file exists
  -o, --out     output directory (default ".")
```

