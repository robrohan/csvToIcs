# CSV to ICS

This is a simple utility to convert a CSV file to an ICS file. I needed something simple to help make training plans into calendar events. This isn't production level software, this isn't a product, it's a quick 'n dirty utility.

It takes a CSV file in the form of:

| Subject | Start Date | Start Time | End Date | End Time | All Day Event | Description | Location | Private |
|---------|------------|------------|----------|----------|---------------|-------------|----------|---------|
| 🛏 Rest | 2018-08-06 |            |          |          |     TRUE      | Things and stuff |      |  FALSE |
| 🛏 Rest | 2018-08-07 |            |          |          |     TRUE      | Things and stuff |      |  FALSE |

or, in a more raw format:

    Subject,Start Date,Start Time,End Date,End Time,All Day Event,Description,Location,Private
    🛏 Rest,2018-08-06,,,,TRUE,,,FALSE

You can create the training plan in Excel, Numbers, scim, or whatever program you have that can export into that specific format.

## Limitations

This will only generate all day events, and it uses the start date. In fact, it only uses the _Subject_, _Start Date_, and _Description_ fields. You can hide the others in your spreadsheet application, but the exported CSV must have all the fields.

The reason the other fields are there is: this particular CSV format is supported by other applications (Google Calendar for example), and it just seemed like a good idea to stick to the format.

Unlike Google Calendar, this application supports using Emojis in the Subject.

## Usage

From a terminal:

    $ ./csvToIcs ./example.csv ./example.ics

## Template

I've added a Numbers file that you can use as a template if you wish