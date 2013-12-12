path2pathAction
========

ver.20131212

Batch make symolic links or copy depneded on json format file define.
developed by Go.

usage: path2pathAction [`-sl`|`-c`] &lt;file_path&gt;

# Json Format

```json
    {
        "case_name": {
            "path_type": "[absolute|abs|a | relative|rel|r]",
            "src_path_parent": "source_path_parent",
            "src_path_children": ["source_child_#1", "source_child_#2", ...],
            "dest_path": "destination_path"
        },
        
        .
        .
        .
        (other cases)
    }
```

- case_name: mark this case, you can define many cases in one json file.
- src_path: define action source. use two layers, so you can make link or copy from many source to destination path.
- dest_path: destination path.

# Example

Consider json content as below, and locate at path `/User/MyComputer/Files/PathsPairTest.json`:

```json
    {
        "Case1": {
            "path_type": "absolute",
            "src_path_parent": "/User/MyComputer/SourceFiles",
            "src_path_children": ["file1.f", "imgae1.png", "folder/fileInFolder.txt"],
            "dest_path": "/User/MyComputer/DestinationFiles"
        },
        
        "Case2": {
            "path_type": "rel",
            "src_path_parent": "SourceFiles",
            "src_path_children": ["file1.f", "imgae1.png", "folder/fileInFolder.txt"],
            "dest_path": "DestinationFiles"
        }
    }
```

and source in `Case1` is:

<pre>
User/MyComputer/SourceFiles/
                        |-file1.f
                        |-imgae1.png
                        |- folder/
                            |-fileInFolder.txt 
</pre>

in `Case2`: 

<pre>
User/MyComputer/Files/
                        |-file1.f
                        |-imgae1.png
                        |- folder/
                            |-fileInFolder.txt 
</pre>

use command `$ path2pathAction -sl /User/MyComputer/Files/Test.json`
then `Case1` would make symbolic links of "file1.f", "imgae1.png" and "fileInFolder.txt" to "/User/MyComputer/DestinationFiles" like below:

<pre>
User/MyComputer/DestinationFiles/SourceFiles
                                    |-file1.f
                                    |-imgae1.png
                                    |- folder/
                                            |-fileInFolder.txt
</pre>

and Case2 do the same action but in json file's relative position:

<pre>
User/MyComputer/Files/DestinationFiles/
                                    |-file1.f
                                    |-imgae1.png
                                    |-fileInFolder.txt                                
</pre>    


(**attention "fileInFolder.txt" not include it's folder**)

# Future Works

- Add unit test.
