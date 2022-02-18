import argparse
from encodings import utf_8
import os
import datetime

# returns the name of the post file in the format 0n-filename.md
def slugify(s):
    if "'" in s:
        s = s.replace("'", "")
    num_of_posts_files = len(os.listdir("./posts"))
    index_to_be_added = "0" + (str(num_of_posts_files % 99))

    return index_to_be_added + "-" + "-".join(s.split(" "))


parser = argparse.ArgumentParser(description="setup blog working files and directories")

parser.add_argument(
    "-n",
    dest="new_post",
    action="store_true",
    required=True,
    help="initiate (n)ew post mode",
)
parser.add_argument("-t", dest="title", action="store", help="post (t)itle")
parser.add_argument(
    "-a", dest="arabic", action="store_true", help="creates an (a)rabic formatted file"
)
parser.add_argument("-l", dest="list", action="store_true", help="(l)ist posts")

args = parser.parse_args()

x = os.listdir(".")
if "index.html" not in x:
    html = open("index.html", "w")
    print("creating html file")
if "styles.css" not in x:
    css = open("styles.css", "w")
    print("creating styles file")
if "posts" not in x:
    os.mkdir("./posts")
    print("creating posts directory")

if args.new_post:
    if args.title:
        slug = slugify(args.title)
        if args.arabic:
            with open(f"./posts/{slug}.rtl.md", "w", encoding="utf_8") as md:
                md.write(f"u{args.title}\n\n")
                md.write(f"#programming\n\n")
                md.write(f"-" * 3 + "\n\n")
        else:
            with open(f"./posts/{slug}.md", "w") as md:
                md.write(f"{args.title}\n\n")
                md.write(f"#programming\n\n")
                md.write(f"-" * 3 + "\n\n")
        print("add new post")
    else:
        print("title needs to be added")

if args.list:
    print(os.listdir("./posts"))
