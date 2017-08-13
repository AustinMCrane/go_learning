-- Schema for api
CREATE TABLE guest (
  id SERIAL,
  last_name char varying(255),
  first_name char varying(255),
  wedding_party boolean default false,
  rsvp boolean default false
);

CREATE TABLE location (
  id SERIAL,
  name char varying(255)
);

CREATE TABLE post (
  id SERIAL,
  body char varying(255),
  image char varying(400)
);

CREATE TABLE timeline (
  id SERIAL,
  body char varying(500),
  time char varying(100),
  image char varying(255)
);
