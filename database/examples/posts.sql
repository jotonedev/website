insert into blog.posts (id, title, description, created_at, updated_at, author_id, content, lang, thumbnail_id)
values  (2301150, 'First Post', 'The first post on jotone.eu', '2023-01-15 12:28:36.554541', '2023-01-15 12:28:36.554541', 1, '# My First post

## What is this?
This is my first website project developed using the Go programming language, so I opted to develop a blog where I can share my various project (*maybe*).
All the website I developed before were or in Python using FastAPI or PHP.

## Why in golang?
I''ve always wanted to learn it, but I just don''t really like languages that are anti-OOP.
I thought that using a sort of functional language for a website was the best option, so I used golang.

## Design
I designed this website to be minimalistic but also feature-complete, such as support for open-graph or other standards.
I also kept a focus on loading speed and responsiveness, preferring svg to raster images where possible.
And also this website supports light-mode and dark-mode produly without js.

## Hosting
This website is hosted in a Docker Swarm with two nodes behind a load balancer proxied with Cloudflare used as CDN.
I choosed Docker Swarm over Kubernets because it was the fastest to setup in my existing environment composed of two nodes.
And I used Cloudflare because they offer a free tier with all the features I needed like CDN, nameserver, R2 and WAF.

## Database
All the posts are inside a free Postgres instance and are stored in markdown to simplify post writing (and also because I prefer markdown to html)', 'en', 1);