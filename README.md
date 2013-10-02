mkslinks
========

ver 0.8

batch make symolic links by json format file
developed by Go

usage: mkslinks <file_path>

define file and source files must be  at the same directiory.

json format:
{
    "files": [<file1>, <file2>, ... ],
    "toPath": <target_parth>
}
