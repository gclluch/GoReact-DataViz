// App.test.js
import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

jest.mock('./TimeSeriesChart', () => () => <div>TimeSeriesChart Mock</div>);

// Verifies that the App component renders correctly and includes two instances of the 
// TimeSeriesChart component, ensuring the main structure of the application is as expected.
test('renders App with TimeSeriesChart components', () => {
  render(<App />);
  const chartComponents = screen.getAllByText('TimeSeriesChart Mock');
  expect(chartComponents).toHaveLength(2); // Expect two instances of the TimeSeriesChart component
});
