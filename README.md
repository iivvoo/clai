# CLAI - Command Line AI

`clai` is a command line integration for openai. It's setup to help you learn
new shell commands and construct more complex commands. E.g.

```
$ clai "group a tab separated file by first column, sum them by third column"
awk -F'\t' '{sum[$1] += $3} END {for(group in sum) print group, sum[group]}' file.txt

This command uses awk to read a tab-separated file, it creates an array 'sum' where the keys are the values in the first column and the values are the sum of the third column for each group. Finally, it prints the group and the sum.
```

```
$ "find the largest file"
To find the largest file in a directory and its subdirectories, you can use the `find` and `du` commands together:

\`\`\`bash
find /path/to/directory -type f -exec du -sh {} + | sort -rh | head -n 1
\`\`\`

This command will find all the files in the specified directory and its subdirectories, calculate their sizes using `du`, sort them in descending order, and display the largest file with its size.
```

By default it assumes a gnu / linux / bash environment, but by explicitly mentioning, e.g. macos or powershell you can get
it to provide solutions for other environments.

```
C:\> clai.exe "find the largest file using powershell"
To find the largest file using PowerShell, you can use the following command:

\`\`\`powershell
Get-ChildItem -File | Sort-Object -Property Length -Descending | Select-Object -First 1
\`\`\`

This command uses the `Get-ChildItem` cmdlet to list all files in the current directory, piped to `Sort-Object` to sort them by file size in descending order, and then piped to `Select-Object` to select just the first (largest) file.
```

## Prerequisites and configuration

Configuration is done through a config.yml file that can be placed in $HOME/.clai/ or $HOME/.config/clai/

You will need an OpenAI api key. Instructions on how to get one are <here>

Here you can also tweak the system prompt (make it more suitable to your needs or environment) and disable the disclaimer once you  get fed up with it and are aware that you should always double check anything generated or suggested by an AI


