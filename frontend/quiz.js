async function submitQuiz() {
    const answers = {};
    const questions = ['city', 'year', 'country', 'president', 'mountain', 'canyon'];

    questions.forEach(question => {
        const selectedOption = document.querySelector(`input[name="${question}"]:checked`);
        if (selectedOption) {
            answers[question] = selectedOption.value;
        }
    });

    try {
        const response = await fetch('http://localhost:8080/submitQuiz', { 
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ answers })
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const result = await response.json();
        document.getElementById('result').innerText = `Your score is: ${result.correctAnswers} out of ${questions.length}`;
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        document.getElementById('result').innerText = 'There was an error submitting your quiz. Please try again.';
    }
}
