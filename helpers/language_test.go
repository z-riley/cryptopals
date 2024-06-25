package helpers

import (
	"testing"
)

func TestRelativeLetterFreqs(t *testing.T) {
	relativeLetterSimilarityEng(`I just returned from the greatest summer vacation! It was so fantastic, I never wanted it to end. I spent eight days in Paris, France. My best friends, Henry and Steve, went with me. We had a beautiful hotel room in the Latin Quarter, and it wasn’t even expensive. We had a balcony with a wonderful view.

We visited many famous tourist places. My favorite was the Louvre, a well-known museum. I was always interested in art, so that was a special treat for me. The museum is so huge, you could spend weeks there. Henry got tired walking around the museum and said “Enough! I need to take a break and rest.”

We took lots of breaks and sat in cafes along the river Seine. The French food we ate was delicious. The wines were tasty, too. Steve’s favorite part of the vacation was the hotel breakfast. He said he would be happy if he could eat croissants like those forever. We had so much fun that we’re already talking about our next vacation!`)
}

func TestIsLikelyEnglish(t *testing.T) {

	for _, tc := range []struct {
		text     string
		expected bool
	}{
		{
			text: `I just returned from the greatest summer vacation! It was so fantastic, I never wanted it to end. I spent eight days in Paris, France. My best friends, Henry and Steve, went with me. We had a beautiful hotel room in the Latin Quarter, and it wasn’t even expensive. We had a balcony with a wonderful view.
We visited many famous tourist places. My favorite was the Louvre, a well-known museum. I was always interested in art, so that was a special treat for me. The museum is so huge, you could spend weeks there. Henry got tired walking around the museum and said “Enough! I need to take a break and rest.”
We took lots of breaks and sat in cafes along the river Seine. The French food we ate was delicious. The wines were tasty, too. Steve’s favorite part of the vacation was the hotel breakfast. He said he would be happy if he could eat croissants like those forever. We had so much fun that we’re already talking about our next vacation!`,
			expected: true,
		},
		{
			text: `Doctor Klein: Good morning, Cecilia, how are you feeling today?
Cecilia: I do not feel very well, Doctor Klein. I hope that you can treat my illness.
Doctor Klein: I’m sorry that you feel very sick. Tell me some of your symptoms so that I can give you a proper diagnosis.
Cecilia: I have not felt well since yesterday afternoon. The symptoms began with a headache and a gradual sore throat. I drank tea with lemon and honey, and I went to bed early. However, I am extremely exhausted, and I don’t feel any better today.
Doctor Klein: I have seen these symptoms recently in some of my other patients. I’ll check your temperature and examine your throat in order to give you a proper diagnosis.
Cecilia: Thank you, Doctor.
Doctor Klein: Oh, my! I can already see that your throat is very red. Your temperature of 100 degrees indicates that you are running a mild fever. I am afraid that you might have the flu.
Cecilia: What is the best way to cure my symptoms?
Doctor Klein: You will need plenty of rest, and you should drink fluids frequently in order to stay hydrated. You’ve also just started to show symptoms, so I can prescribe you a medication that can reduce fever and shorten the duration of your illness.
Cecilia: Should I stay home from work as well?
Doctor Klein: Yes, you should remain in bed until the fever breaks. You should also wait until 24 hours after the fever has broken before you return to work. You do not want to risk getting your coworkers sick as well.
Cecilia: I suppose I will just take it easy and relax for a couple of days. Thank you, doctor, for all of your help!
Doctor Klein: No problem! Try your best to rest for a couple of days. I hope you feel better soon!`,
			expected: true,
		},
		{
			text:     `p/-9i/50234gfrdeafej	23`,
			expected: false,
		},
		{
			text:     `I'm a short sentence in English`,
			expected: true,
		},
		{
			text:     `それはいい考えですね`,
			expected: false,
		},
	} {
		actual := IsLikelyEnglish(tc.text)
		if tc.expected != actual {
			t.Errorf("Expected: %t\nGot: %t\nText: %s", tc.expected, actual, tc.text)
		}
	}

}
