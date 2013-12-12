path2pathAction
========

ver 10.

batch make symolic links or copy by json format file define.
developed by Go.

usage: path2pathAction [-sl | -c] &lt;file_path&gt;

define "file" and "source files" must be at THE SAME DIRECTIORY.

```json
    json format:
    {
        "files": ["file1", "file2", ... ],
        "toPath": &lt;target_path&gt;
    }
```