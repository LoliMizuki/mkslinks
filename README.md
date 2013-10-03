mkslinks
========

ver 0.9

batch make symolic links by json format file
developed by Go

usage: mkslinks &lt;file_path&gt;

define "file" and "source files" must be at THE SAME DIRECTIORY.

json format:<br />
{<br />
&nbsp;&nbsp;&nbsp;&nbsp;"files": ["file1", "file2", ... ],<br />
&nbsp;&nbsp;&nbsp;&nbsp;"toPath": &lt;target_path&gt;<br />
}<br />
