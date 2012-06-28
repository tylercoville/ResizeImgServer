*This is a Real World Problem. Please read about Real World Problems at
the bottom before attempting.*

For the first CodeSprint, we at InterviewStreet manually grabbed the
logos straight off of websites and resized them. This time around, we
actually have company profiles, so participating companies can upload
their logos.

Unfortunately, we didn’t resize logos on upload, so when it came time to
release the wall of participating companies, we got results like this:

![][0]

Now, it’s not particularly difficult to batch resize logos, but not all
logos are super-convenient transparent .pngs; some of them have solid
backgrounds. As a result, resizing them makes them a little too small:

![][1]

**Task: **Create a webservice that will resize logos with backgrounds by
logo size, not image size.

**Input:**

-   An image
-   Width and Height parameters (in pixels)
-   Padding (in pixels)

**Output:**

The resized logo, with the logo portion fitting just inside the padding
of the image. (Don’t stretch to fit both width and height. Maintain the
logo aspect ratio)

Please note that if the original logo came with a colored background,
the entire background of the image should be that color. (In other
words, you may need to extend the background to cover what would
normally be white borders in the image) Transparent backgrounds are kept
transparent.


**Sample Input:**
http://ec2-AA-BB-CC-DD.compute-1.amazonaws.com/logoresize/?image=http://example.com/img.png&width=200&height=100&padding=20

![][2]

**Sample Output:**

Return the resized logo, with the same filename and filetype as the
input

![][3]

**Constraints:**

Width <= 640px

Height <= 480px

Padding <= 25% of min(width, height)

Image filetype is restricted to jpg and png

Logos are defined by the smallest inscribed rectangle such that
everything outside that rectangle is a single color.

The logo should be centered both horizontally and vertically (the
inscribed rectangle, that is. Visually, that may or may not look
centered)

Image backgrounds are a single solid color; you won’t be given
background patterns or gradients

**Scoring:**

We will test against multiple images, starting with transparent,
to-the-edge logos, with the hardest image being a logo with a background
requiring additional fill (like the MemSQL logo example)

Because image resizing programs are slightly fuzzy, we will consider a
test case passed if 80-95% of the background color pixels match (the
exact percentage depends on the complexity of the logo).

In general though, our hypothesis is that Real World Problems are much
more about your code and approach than your score.

**Real World Problem Instructions:**

Please write your code on your local machine. Design for a standard
ubuntu Amazon Machine Instance, and you are free to write in any
language that runs on that machine. When you’re ready to test and submit
*both* AMI Real World
Problems, click the Participate button below, enter your [public
key][] to check out an instance for 1 hour and load your code via SSH.

After one hour, we will automatically shut down your instance and save
your /var/www folder. Unfortunately, we weren’t able to develop
real-time checking of this problem in time, so we will be grading this
problem after the contest is over (but before companies will see your
profiles). *Please include a Install.txt readme* so we can load up any
additional packages and libraries that you used.

If you have no idea what any of the above means (and you’re still
thinking of trying the problem) check out this hastily written FAQ
and join in our chatroom. We’ll help you as much as we ca TRUNCATED!
Please download pandoc if you want to convert large files.

  [0]: https://static.interviewstreet.com/recruit/questions/5138/image_1.png
  [1]: https://static.interviewstreet.com/recruit/questions/5138/image_2.png
  [2]: https://static.interviewstreet.com/recruit/questions/5138/image_3.png
  [3]: https://static.interviewstreet.com/recruit/questions/5138/image_4.png