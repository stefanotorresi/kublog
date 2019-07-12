A Kubernetes blog
=================

This exercise consists in writing a blog with the typical blog
resources created as custom resource definitions (CRD's). Typically, a
blog will consist at least:

* Blog posts
* Comments
* Comment upvotes (+1)

Relationship between the three entities is as follows:
  Blog post < 1 --- N > Comments < 1 --- N > Comment upvotes

Feel free to add the attributes to each resource that you think make
sense (e.g. a blog post should probably contain a title, a date and a body).

The main task of the controller or controllers that you will write is
to ensure that the status of the Blog posts and Comments reflect the
following:

* A blog post status must have an attribute `CommentCount`; whenever a
  comment is created or deleted, your controller will have to update
  the `CommentCount` on the blog post status.

* A blog post comment status must have an attribute `UpvoteCount`; whenever
  an upvote is created or deleted, your controller will have to update
  the `UpvoteCount` on the blog post comment status.

Note: it's not required for the controllers to run inside the cluster,
as long as they perform the desired logic.