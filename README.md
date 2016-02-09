# mars

Last year by far the best movie I saw at the cinema was "The Martian", known in China as "Mars Rescue" and shown a few months after its release in the US. After looking around at some calendaring systems for Mars, I found none of them suitable so I invented one myself.

The concepts of "Martian day" and "Martian year" are quite intuitive, and form the basis of this calendar.

## Intra-day

The system used by US robotic missions to Mars of making the Martian second 2.75% longer than the Earth second is a good system, so there's 60 Martian seconds in a Martian minute, 60 of those in a Martian hour, and 24 of those in a Martian day, called a "sol". The Martian time zone used by US missions is whatever puts the Sun directly overhead at Martian noon, but perhaps later on Mars should be divided into 24 time zones.

## Week

The 7-day week has been around for at least 2,500 years, and is now standardized by virtually every society on Earth. A 7-sol week would also be a natural choice for Mars, but there's one additional feature of a Martian week which would make it special, and forms the basis for the subsequent decisions in designing this calendaring system: *The day of the week on Mars should be the same as the current day of week on Earth*.

Because 1 Martian "sol" is the same as 1 day and 40 minutes on Earth, matching the Martian day of week to the Earth one will cause a 1-sol slip around every 37 sols, so a day of week will need to be skipped to reset the Martian day of week back to the Earth one. A second design decision is that *the same day of week should be skipped each time*. Because the present international standard is that Monday is the first day of the week and Sunday the seventh, perhaps it's best that the last day of the week, Sunday, be skipped. This would happen once every 5 or 6 weeks.

## Month

There should be a time period roughly matching what's known as a month on Earth. The third design decision is that *a Martian month would end whenever a Sunday is skipped*. So a Martian month would always be 34 or 41 days long. On average, every third month would be 41 days long, and the other two 34 days long. Because a Martian year is about 668 sols long, this puts the number of Martian months per Martian year as usually 18 and sometimes 19.

The timing and name of a Martian month could roughly follow those on Earth, or they could follow the seasons on Mars, or perhaps a bit of both. We've tried to do both in our fourth design decision: *every second Martian month name will match the current month on Earth, and every other Martian month will match the seasons on Mars*. Tentatively, we've used the English month names January to December for the Earth months, and the Horoscope signs Aries to Pisces for the seasons on Mars. The same Earth month name won't be used twice only one Earth year apart but forward to the following Earth month if necessary. Therefore, all month names in a single Martian year will be unique.

## Year

I've written a quick script in Golang to look at the distribution of month names using these rules over 25 Martian years. Aries marks the beginning of Northern Spring, and the distorted season lengths on Mars are accounted for...

```
Jan:22, Feb:16, Mar:16, Apr:19, May:20, Jun:21, Jul:19, Aug:19, Sep:20, Oct:18, Nov:18, Dec:22, 
Ari:21, Tau:21, Gem:23, Can:20, Leo:21, Vir:17, Lib:15, Sco:18, Sag:15, Cap:16, Aqu:21, Pis:22, 
```

The shortest seasonal month names are Libra and Saggitarius, which matches winter and autumn being the shortest months on Mars. February, the shortest month on Earth, is also among the shortest Earth-month-named months on Mars.

The full data, where a `+` after the month name denotes a 41-day month...

```
月:    1     2     3     4     5     6     7     8     9    10    11    12    13    14    15    16    17    18    19
 1年: Jan+  Ari   Mar+  Tau   Jun   Gem+  Aug   Can   Nov+  Leo   Feb   Lib+  Apr   Sco   Jul+  Cap   Sep   Aqu+  Dec
 2年: Pis   Jan+  Tau   May   Gem+  Jun   Can   Oct+  Leo   Nov   Lib+  Feb   Sco   Apr+  Cap   Jul   Aqu+  Sep
 3年: Pis   Dec+  Ari   Mar   Gem+  May   Can   Aug+  Leo   Oct   Vir+  Jan   Sco   Feb+  Sag   Jun   Aqu+  Jul   Pis+
 4年: Nov   Ari   Dec+  Tau   Mar   Gem+  May   Leo   Aug+  Vir   Oct   Sco+  Jan   Sag   Apr+  Aqu   Jun   Pis+
 5年: Sep   Ari   Nov+  Tau   Dec   Gem+  Mar   Can   Jul+  Vir   Aug   Lib+  Oct   Sag   Jan+  Cap   Apr   Pis+
 6年: Jun   Ari   Sep+  Tau   Nov   Gem+  Feb   Can   May+  Leo   Jul   Lib+  Oct   Sco   Dec+  Cap   Jan   Aqu+  Apr
 7年: Ari+  Jun   Tau   Sep+  Gem   Nov   Can+  Feb   Leo   May+  Lib   Jul   Sco+  Oct   Cap   Dec+  Aqu   Mar
 8年: Pis+  Jun   Ari   Aug+  Gem   Sep   Can+  Jan   Leo   Feb+  Vir   May   Sco+  Jul   Sag   Oct+  Aqu   Dec   Pis+
 9年: Mar   Ari   Jun+  Tau   Aug   Can+  Nov   Leo   Jan+  Vir   Apr   Sco+  May   Sag   Sep+  Aqu   Oct   Pis+
10年: Dec   Ari+  Mar   Tau   Jun+  Gem   Aug   Leo+  Nov   Vir   Jan+  Lib   Apr   Sag+  Jul   Cap   Sep+  Pis
11年: Dec   Ari+  Feb   Tau   May+  Gem   Jun   Can+  Oct   Vir   Nov+  Lib   Jan   Sag+  Apr   Cap   Jul+  Aqu   Sep
12年: Ari+  Dec   Tau   Feb+  Gem   May   Can+  Aug   Leo   Oct+  Lib   Jan   Sco+  Mar   Cap   Jun+  Aqu   Jul
13年: Pis+  Sep   Tau+  Dec   Gem   Apr+  Can   May   Leo+  Aug   Vir   Oct+  Sco   Jan   Cap+  Mar   Aqu   Jun+
14年: Pis   Jul   Ari+  Nov   Tau   Dec+  Can   Apr   Leo+  May   Vir   Aug+  Sco   Oct   Sag+  Jan   Aqu   Mar+  Pis
15年: Jun   Ari+  Sep   Tau   Nov+  Gem   Feb   Leo+  Apr   Vir   Jul+  Lib   Aug   Sag+  Dec   Cap   Jan+  Pis
16年: May   Ari+  Jun   Tau+  Sep   Gem   Nov+  Can   Feb   Vir+  Apr   Lib   Jul+  Sag   Oct   Cap+  Dec   Aqu   Mar+
17年: Ari   May   Tau+  Aug   Gem   Sep+  Can   Jan   Leo+  Feb   Lib   Jun+  Sco   Jul   Cap+  Oct   Aqu   Dec+
18年: Pis   Mar   Tau+  May   Gem   Aug+  Can   Nov   Leo+  Jan   Vir   Apr+  Sco   Jun   Cap+  Sep   Aqu   Oct+
19年: Pis   Dec   Ari+  Mar   Gem+  May   Can   Aug+  Leo   Nov   Vir+  Jan   Sco   Apr+  Sag   Jun   Aqu+  Sep   Pis
20年: Dec+  Ari   Feb   Tau+  May   Gem   Jul+  Leo   Aug   Vir+  Nov   Lib   Jan+  Sag   Apr   Aqu+  Jun   Pis
21年: Sep+  Ari   Dec   Tau+  Feb   Gem   May+  Can   Jul   Vir+  Oct   Lib   Jan+  Sag   Mar   Cap+  Apr   Pis
22年: Aug+  Ari   Sep   Tau+  Dec   Gem+  Feb   Can   May+  Leo   Jul   Lib+  Oct   Sco   Jan+  Cap   Mar   Aqu+  Jun
23年: Pis   Aug+  Tau   Nov   Gem+  Dec   Can   Apr+  Leo   May   Lib+  Sep   Sco   Oct+  Cap   Jan   Aqu+  Mar
24年: Pis   Jun+  Ari   Aug   Gem+  Nov   Can   Feb+  Leo   Apr   Vir+  Jul   Sco   Sep+  Sag   Dec   Aqu+  Jan   Pis
25年: May+  Ari   Jun   Tau+  Aug   Gem+  Nov   Leo   Feb+  Vir   Apr   Sco+  Jul   Sag   Sep+  Aqu   Dec   Pis+
```

## Summary of Rules

* The day of the week on Mars should be the same as the current day of week on Earth

* The same day of week should be skipped each time

* A Martian month would end whenever a Sunday is skipped

* Every second Martian month name will match the current month on Earth, and every other Martian month will match the seasons on Mars

Of course there's many more rules we could add to fine-tune the idea. When the Earth is closest to Mars, we could have two consecutive Earth-month-named Martian months, and when Earth is furthest away, have two consecutive season-named Martian months. But these rules roughly show how the 24 month names would be distributed. And this is assuming there's no flaws in my program.

## License

Copyright © 2016 Gavin "Groovy" Grover

Distributed under the same BSD-style license as Go that can be found in the LICENSE file.

